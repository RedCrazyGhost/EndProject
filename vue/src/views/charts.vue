<template>
	<div class="container">
		<el-row :gutter="24">
			<el-col :span="5"><el-autocomplete
					v-model="useTable"
					:fetch-suggestions="querySearchAsync"
					placeholder="请选择分析的表"
					@select="handleSelect"
			/></el-col>
			<el-col :span="5">
			<el-input v-model="balanceCol"
								class="w-50 m-2" placeholder="余额列"/>
			</el-col>
			<el-col :span="5">
				<el-input v-model="amountCol"
									class="w-50 m-2" placeholder="交易金额"/>
			</el-col>
			<el-col :span="5">
				<el-input v-model="grouCol"
									class="w-50 m-2" placeholder="聚合列"/>
			</el-col>
			<el-button type="primary"    @click="checkTable">分析</el-button>
		</el-row>
		<div v-if="checkTableData.length!=0">
		<div style="text-align: center;color: red">异常数据</div>
		<el-table :data="checkTableData" border class="table" header-cell-class-name="table-header">
			<el-table-column v-for="(value,key) in checkTableData[0]" :prop="key" :label="useTableHead.get(key)" align="center"></el-table-column>
		</el-table>
		</div>
		<div class="schart-box" >
			<schart class="schart" canvasId="bar" :options="options1"></schart>
		</div>
		<div class="schart-box"  >
			<schart class="schart" canvasId="line" :options="options2"></schart>
		</div>
		<div class="schart-box" >
			<schart class="schart" canvasId="pie" :options="options3"></schart>
		</div>
		<div class="schart-box" >
			<schart class="schart" canvasId="ring" :options="options4"></schart>
		</div>
		</div>
</template>

<script setup lang="ts" name="basecharts">
import {FormInstance} from "element-plus";

let UserId=localStorage.getItem("UserId")

import Schart from 'vue-schart';
import {onMounted, ref} from "vue";
import axios from "axios";
import object from "async-validator/dist-types/validator/object";
import app from "../App.vue";

const useTable = ref('')

interface TableItem {
	value: string
}

const tables = ref<TableItem[]>([])

let timeout: NodeJS.Timeout
const querySearchAsync = (queryString: string, cb: (arg: any) => void) => {
	const results = queryString
			? tables.value.filter(createFilter(queryString))
			: tables.value

	clearTimeout(timeout)
	timeout = setTimeout(() => {
		cb(results)
	}, 3000 * Math.random())
}
const createFilter = (queryString: string) => {
	return (restaurant: TableItem) => {
		return (
				restaurant.value.toLowerCase().indexOf(queryString.toLowerCase()) === 0
		)
	}
}

const handleSelect = (item: TableItem) => {
	console.log(item)
	console.log(useTable)
}

const useTableHead =new Map()
function useTableHeadConf() {
	useTableHead.set("checkResult", "检测结果")
	useTableHead.set("rowNum", "行数")
	useTableHead.set("failResult", "错误结果")
}

onMounted(async () => {
	await axios.get("http://localhost:8080/data/list?UserId="+UserId).then(function (r) {
		let data=r.data.data

		for(let i=0;i<data.length;i++){
			let obj=data[i];
				obj["value"]=obj["TABLE_NAME"]
				delete  obj["TABLE_NAME"]
		}
		tables.value = data
	})
})


const checkTableData = ref<Object>([]);
let grouCol=ref();
let balanceCol=ref();
let amountCol=ref();
let options1=ref({})
let options2=ref({})
let options3=ref({})
let options4=ref({})
async function checkTable() {
	await axios.get("http://localhost:8080/data/head?UserId="+UserId+"&TableName="+useTable.value).then(function (r) {
		if (r.data.data!=null){
			useTableHeadConf()

			for(const item of r.data.data){
				useTableHead.set(item["col_name"],item["col_comment"])
			}
		}else{
			useTableHead.clear()
			useTableHeadConf()
		}

	})
	// 获取表格数据
	if (balanceCol.value!=undefined&&amountCol.value!=undefined) {
		await axios.get("http://localhost:8080/data/check?UserId=" + UserId + "&TableName=" + useTable.value + "&colBalance=" + balanceCol.value + "&colAmount=" + amountCol.value).then(function (r) {
			if (r.data.data != null) {
				checkTableData.value = r.data.data
			} else {
				checkTableData.value = []
			}

		})
	}

	// 聚合
	if (grouCol.value!=undefined) {
		await axios.get("http://localhost:8080/data/group?UserId=" + UserId + "&TableName=" + useTable.value + "&GroupCol=" + grouCol.value).then(function (r) {
			if (r.data.data != null) {
				let labels = []
				let data = []
				for (let item of r.data.data) {
					labels.push(item.k)
					data.push(item.v)
				}
				options3.value = {
					type: 'pie',
					title: {
						text: '银行交易类型数量图'
					},
					legend: {
						position: 'left'
					},
					bgColor: '#fbfbfb',
					labels: labels,
					datasets: [
						{
							data: data
						}
					]
				};
			} else {
			}

		})
	}




	options1.value = {
		type: 'bar',
		title: {
			text: '银行流水交易数据安全图'
		},
		bgColor: '#fbfbfb',
		labels: ['状态'],
		datasets: [
			{
				data: [134]
			},{
				fillColor: 'rgba(241, 49, 74, 0.5)',
				data:[3]
			}
		]
	};
	options2.value = {
		type: 'line',
		title: {
			text: '银行交易流水趋势图'
		},
		bgColor: '#fbfbfb',
		labels: ['7月',
			'8月',
			'9月',
			'10月',
			'11月',
			'12月',
			'1月',
		],
		datasets: [
			{
				data: [123581591.72,
					123596792.12,
					123588812.62,
					123590343.27,
					123819159.74,
					124397544.27,
					124371088.42
				]
			}
		]
	};

	options4.value = {
		type: 'ring',
		title: {
			text: '单笔交易范围'
		},
		showValue: false,
		legend: {
			position: 'bottom',
			bottom: 40
		},
		bgColor: '#fbfbfb',
		labels: ['<=100', '<=1000', '>1000'],
		datasets: [
			{
				data: [108, 2, 27]
			}
		]
	};
}

</script>

<style scoped>
.schart-box {
	display: inline-block;
	margin: 20px;
}
.schart {
	width: 600px;
	height: 400px;
}
.content-title {
	clear: both;
	font-weight: 400;
	line-height: 50px;
	margin: 10px 0;
	font-size: 22px;
	color: #1f2f3d;
}
</style>
