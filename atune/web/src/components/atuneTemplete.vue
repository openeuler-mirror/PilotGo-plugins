<template>
    <div>
        <el-form :model="form" class="custom-form">
            <el-form-item label="调优对象">
                <el-input v-model="form.tuneName"></el-input>
            </el-form-item>
            <el-form-item label="工作目录">
                <el-input v-model="form.workDir"></el-input>
            </el-form-item>
            <el-form-item label="环境准备">
                <el-input v-model="form.prepare"></el-input>
            </el-form-item>
            <el-form-item label="智能调优">
                <el-input v-model="form.tune"></el-input>
            </el-form-item>
            <el-form-item label="环境恢复">
                <el-input v-model="form.restore"></el-input>
            </el-form-item>
            <el-form-item label="注意事项">
                <el-input v-model="form.note" class="custom-input" type="text" :rows="4"></el-input>
            </el-form-item>
        </el-form>
        <el-form class="centered-buttons">
            <el-button type="primary" @click="onSubmit" class="custom-button">保存</el-button>
            <el-button class="custom-button">取消</el-button>
        </el-form>
    </div>
</template>

<script lang='ts' setup>
import { ref, onUpdated, reactive } from 'vue'
import { getAtuneInfo } from '@/api/atune'

let props = defineProps({
    selectedNodeData: {
        type: String,
        default: ""
    }
})

const atuneName = ref(props.selectedNodeData)

onUpdated(() => {
    atuneName.value = props.selectedNodeData
    fetchAtuneInfo();
})

const form = reactive({
    tuneName: "",
    workDir: "",
    prepare: "",
    tune: "",
    restore: "",
    note: ""
})
const fetchAtuneInfo = () => {
    if (atuneName.value) {
        getAtuneInfo({ name: atuneName.value })
            .then((res) => {
                if (res.data && res.data.code === 200) {
                    const data = res.data.data;
                    // 判断数据结构类型
                    if (data.BaseTune) {
                        // 如果有 BaseTune
                        form.tuneName = data.BaseTune.tuneName || "";
                        form.workDir = data.BaseTune.workDir || "";
                        form.prepare = data.BaseTune.prepare || "";
                        form.tune = data.BaseTune.tune || "";
                        form.restore = data.BaseTune.restore || "";
                    } else {
                        form.tuneName = data.tuneName || "";
                        form.workDir = data.workDir || "";
                        form.prepare = data.prepare || "";
                        form.tune = data.tune || "";
                        form.restore = data.restore || "";
                    }
                    form.note = data.note || "";
                    console.log('获取到的调优信息：', data);
                } else {
                    console.log('获取调优信息时出错:', res.data.msg)
                }
            })
    } else {
        console.warn('atuneName 为空，无法获取调优信息');
    }
}
const onSubmit = () => {
    console.log('submit!')
}
</script>

<style lang = 'less' scoped>
.custom-form {
    margin-left: 25px;
    margin-right: 20px;
    outline-width: 120px;

    .custom-input {
        white-space: pre-wrap;
        text-align: left; // 靠左
        vertical-align: top; // 靠上
        height: 150px;
    }
}


.centered-buttons {
    display: flex;
    justify-content: center;
    margin-top: 20px;

    .custom-button {
        margin-right: 7px;
    }
}
</style>
