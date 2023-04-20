<template>
	<div class="container">
		<div class="schart-box">
			<div class="content-title">柱状图</div>
			<schart class="schart" canvasId="bar" :options="options1"></schart>
		</div>
		<div class="schart-box">
			<div class="content-title">折线图</div>
			<schart class="schart" canvasId="line" :options="options2"></schart>
		</div>
		<div>异常数据</div>
		<el-table :data="tableData" border class="table" header-cell-class-name="table-header">
			<el-table-column prop="col_0" label="序号" width="55" align="center"></el-table-column>
			<el-table-column prop="col_1" label="摘要"></el-table-column>
			<el-table-column prop="col_2" label="币别"></el-table-column>
			<el-table-column prop="col_3" label="钞汇"></el-table-column>
			<el-table-column prop="col_4" label="交易日期"></el-table-column>
			<el-table-column prop="col_5" label="交易金额"></el-table-column>
			<el-table-column prop="col_6" label="账户余额"></el-table-column>
			<el-table-column prop="col_7" label="交易地点/附言"></el-table-column>
			<el-table-column prop="col_8" label="对方账号与用户名"></el-table-column>
		</el-table>
		<div class="schart-box">
			<div class="content-title">饼状图</div>
			<schart class="schart" canvasId="pie" :options="options3"></schart>
		</div>
		<div class="schart-box">
			<div class="content-title">环形图</div>
			<schart class="schart" canvasId="ring" :options="options4"></schart>
		</div>
	</div>
</template>

<script setup lang="ts" name="basecharts">
import Schart from 'vue-schart';
import {ref} from "vue";

const tableData = ref<TableItem[]>([]);
// 获取表格数据
const getData = () => {
	tableData.value =[
		{
			"col_0": "112",
			"col_1": "消费",
			"col_2": "人民币元",
			"col_3": "钞",
			"col_4": "20221230",
			"col_5": "-20",
			"col_6": "124197551.27",
			"col_7": "支*****店",
			"col_8": "1*****店"
		}, {
			"col_0": "8",
			"col_1": "消费",
			"col_2": "人民币元",
			"col_3": "钞",
			"col_4": "20220806",
			"col_5": "-300",
			"col_6": "123583192.12",
			"col_7": "支*****店",
			"col_8": "1*****店"
		},
		{
			"col_0": "134",
			"col_1": "消费",
			"col_2": "人民币元",
			"col_3": "钞",
			"col_4": "20230111",
			"col_5": "-8000.00",
			"col_6": "124263086.42",
			"col_7": "财*****账",
			"col_8": "1*****账"
		}
	]
};
getData();

const options1 = {
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
const options2 = {
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
const options3 = {
	type: 'pie',
	title: {
		text: '银行交易类型数量图'
	},
	legend: {
		position: 'left'
	},
	bgColor: '#fbfbfb',
	labels: ['消费', '支付机构提现', '电子汇入', '短信服务费', 'ATM存款', '充值', '利息存入','银联入账',
		'跨行转出',
		'冲正',
		'代收付',
		'消费退货',
		'转账支取'
	],
	datasets: [
		{
			data: [89,
				1,
				11,
				6,
				2,
				2,
				2,
				11,
				7,
				3,
				1,
				1,
				1,
			]
		}
	]
};
const options4 = {
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
