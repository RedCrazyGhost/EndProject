<template>
    <div>
        <div class="container">
            <div class="handle-box">
                <el-upload
                    ref="uploadRef"
                    :action="'http://localhost:8080/data/upload?UserId='+UserId"
                    :limit="1"
                    accept=".xlsx,.xls"
                    :auto-upload="false"
                    :on-change="handleChange"
                >
                    <el-button class="mr10" type="primary">预览Excel文件</el-button>
                </el-upload>
                <el-button type="success" @click="submitUpload">上传文件到服务器</el-button>
            </div>
            <el-table :data="showTableData" border class="table" header-cell-class-name="table-header">
                <el-table-column v-for="(val,key) in showTableData[0]" :prop="key" :label="showTableHead.get(key)" align="center"></el-table-column>
            </el-table>
        </div>
    </div>
</template>

<script setup lang="ts" name="import">
let UserId=localStorage.getItem("UserId")
import {UploadFile, UploadFiles, UploadInstance, UploadProgressEvent, UploadProps, UploadUserFile} from 'element-plus';
import { ref, reactive } from 'vue';
import * as XLSX from 'xlsx';
import {log} from "echarts/types/src/util/log";
import {read, readFile} from "xlsx";
const uploadRef = ref<UploadInstance>()
const submitUpload = () => {
    uploadRef.value!.submit()
}
const showTableData =ref<Object>([])
let showTableHead=new Map()
const handleChange = (uploadFile: UploadFile, uploadFiles: UploadFiles) => {
    showTableData.value=[]
    showTableHead=new Map()
    const blob=new Blob([uploadFile.raw as BlobPart])
    const reader = new FileReader();
    reader.readAsArrayBuffer(blob);
    reader.onload = async function (loadEvent: any) {
        const arrayBuffer = loadEvent.target["result"];
        const workbook = await XLSX.read(new Uint8Array(arrayBuffer), {
            type: "array"
        });
        let meta = workbook.Sheets.Sheet1
        let range = workbook.Sheets.Sheet1["!ref"]?.split(":");
        for (let i = range[0].charCodeAt(0); i <= range[1].charCodeAt(0); i++) {
            showTableHead.set(String.fromCharCode(i), meta[String.fromCharCode(i) + "1"].v)
        }
        var regExp = new RegExp('\\d+\\.?\\d*');
        for (let i = 2; i < Number(regExp.exec(range[1])[0]); i++) {
            let obj = {}
            for (let [key, value] of showTableHead) {
                obj[key] = meta[key + i].v
            }
            showTableData.value.push(obj)
        }
    };
};



</script>

<style scoped>
.handle-box {
    display: flex;
    margin-bottom: 20px;
}

.table {
    width: 100%;
    font-size: 14px;
}
.mr10 {
    margin-right: 10px;
}
</style>
