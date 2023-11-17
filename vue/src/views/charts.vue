<template>
	<div class="container">
		<el-row :gutter="24">
			<el-col :span="4"><el-autocomplete
					v-model="useTable"
					:fetch-suggestions="querySearchAsync"
					placeholder="请选择分析的表"
					@select="handleSelect"
			/></el-col>
<!--			<el-col :span="4">-->
<!--				<el-select v-model="datetimeCol" clearable placeholder="请选择时间列">-->
<!--					<el-option v-for="item in colsOptions" :key="'item.value'-1" :label="item.label" :value="item.value"/>-->
<!--				</el-select>-->
<!--			</el-col>-->
			<el-col :span="4">
				<el-select v-model="balanceCol" clearable placeholder="请选择余额列">
					<el-option v-for="item in colsOptions" :key="'item.value'-2" :label="item.label" :value="item.value"/>
				</el-select>
			</el-col>
			<el-col :span="4">
				<el-select v-model="amountCol" clearable placeholder="请选择交易金额列">
					<el-option v-for="item in colsOptions" :key="'item.value'-3" :label="item.label" :value="item.value"/>
				</el-select>
			</el-col>
			<el-col :span="4">
				<el-select v-model="grouCol" clearable placeholder="请选择聚合列">
					<el-option v-for="item in colsOptions" :key="'item.value'-4" :label="item.label" :value="item.value"/>
				</el-select>
			</el-col>
			<el-button type="primary"    @click="checkTable">分析</el-button>
		</el-row>
		<div v-if="useTableHead">
			<div style="text-align: center">表头信息</div>
			<el-table :data="useTableHead" border class="table" header-cell-class-name="table-header">
				<el-table-column v-for="(val,key,index) of useTableHead[0]" :prop="key" :label="key" />
			</el-table>
		</div>
		<div v-if="checkTableData.length!=0">
		<div style="text-align: center;color: red">异常数据</div>
		<el-table :data="checkTableData" border class="table" header-cell-class-name="table-header">
			<el-table-column v-for="(value,key) in checkTableData[0]" :prop="key" :label="useCheckTableHead.get(key)" align="center"></el-table-column>
		</el-table>
		</div>
		<div v-if="options1.type!=undefined" class="schart-box" >
			<schart class="schart" canvasId="pie-1"  :options="options1"></schart>
		</div>
		<div v-if="options3.type!=undefined"  class="schart-box" >
			<schart class="schart" canvasId="pie-2" :options="options3"></schart>
		</div>
		</div>
</template>

<script setup lang="ts" name="basecharts">

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

let useTableHead =ref()
const handleSelect = async (item: TableItem) => {
	await axios.get("http://localhost:8080/data/head?UserId=" + UserId + "&TableName=" + useTable.value).then(function (r) {
		useTableHead.value=[]
		colsOptions.value =[]
		if (r.data.data != null) {
			let obj:any={}
			for(let item of r.data.data){
				obj[item["col_comment"]]=item["col_name"]
				colsOptions.value.push({value:item["col_comment"],label:item["col_comment"]})
			}
			useTableHead.value.push(obj)
			console.log(colsOptions)
		}
	})
}


const useCheckTableHead =new Map()


function useTableHeadConf() {
	useCheckTableHead.set("checkResult", "检测结果")
	useCheckTableHead.set("rowNum", "行号")
	useCheckTableHead.set("failResult", "错误结果")
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
let datetimeCol=ref();
let grouCol=ref();
let balanceCol=ref();
let amountCol=ref();
let colsOptions=ref([])
let options1=ref({})
let options2=ref({})
let options3=ref({})
let options4=ref({})
async function checkTable() {
	await axios.get("http://localhost:8080/data/head?UserId="+UserId+"&TableName="+useTable.value).then(function (r) {
		if (r.data.data!=null){
			useTableHeadConf()

			for(const item of r.data.data){
				useCheckTableHead.set(item["col_name"],item["col_comment"])
			}
		}else{
			useCheckTableHead.clear()
			useTableHeadConf()
		}

	})
	let countTable=0
	//获取数据数量
	await axios.get("http://localhost:8080/data/count?UserId="+UserId+"&TableName="+useTable.value).then(function (r) {
		if (r.data.data != null) {
			countTable=r.data.data.count
		} else {
			countTable=0
		}

	})
	// 获取表格数据
	if (balanceCol.value!=undefined&&amountCol.value!=undefined) {
		await axios.get("http://localhost:8080/data/check?UserId=" + UserId + "&TableName=" + useTable.value + "&colBalance=" + useTableHead.value[0][balanceCol.value] + "&colAmount=" + useTableHead.value[0][amountCol.value]).then(function (r) {
			if (r.data.data != null) {
				checkTableData.value = r.data.data

			} else {
				checkTableData.value = []
			}
			options1.value = {
				type: 'pie',
				title: {
					text: '数据健康饼状图'
				},
				bgColor: '#fbfbfb',
				labels: ['正常','异常'],
				datasets: [
					{
						data: [countTable-checkTableData.value.length,checkTableData.value.length]
					}
				]
			};
		})
	}

	// 聚合
	if (grouCol.value!=undefined) {
		await axios.get("http://localhost:8080/data/group?UserId=" + UserId + "&TableName=" + useTable.value + "&GroupCol=" + useTableHead.value[0][grouCol.value]).then(function (r) {
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
						text: '聚类饼状图'
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
