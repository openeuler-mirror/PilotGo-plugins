<template>
    <div>
        <p>{{ selectedNodeData }}</p>

        <div>获取调优信息</div>
        <button @click="fetchAtuneInfo">获取调优信息</button>

        <el-form :model="form" label-width="120px">
            <el-form-item label="调优对象">
                <el-input v-model="form.tuneName"></el-input>
            </el-form-item>
            <el-form-item label="工作目录">
                <el-input v-model="form.workDir"></el-input>
            </el-form-item>
            <el-form-item label="环境准备">
                <el-input v-model="form.perpare"></el-input>
            </el-form-item>
            <el-form-item label="智能调优">
                <el-input v-model="form.tune"></el-input>
            </el-form-item>
            <el-form-item label="环境恢复">
                <el-input v-model="form.restore"></el-input>
            </el-form-item>
            <el-form-item label="注意事项">
                <el-input v-model="form.note"></el-input>
            </el-form-item>
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
// 使用 onUpdated 钩子来监视 selectedNodeData 的变化
onUpdated(() => {
    atuneName.value = props.selectedNodeData
})

const atuneInfo = ref(null)
const form = reactive({
    tuneName: "",
    workDir: "",
    perpare: "",
    tune: "",
    restore: "",
    note: ""
})
const fetchAtuneInfo = () => {
    if (atuneName.value) {
        // 获取单个模板信息
        getAtuneInfo({ name: atuneName.value })
            .then((res) => {
                if (res.data && res.data.code === 200) {
                    atuneInfo.value = res.data;
                    console.log('获取到的调优信息：', atuneInfo.value);
                } else {
                    console.log('获取调优信息时出错:', res.data.msg)
                }
            })
    } else {
        console.warn('atuneName 为空，无法获取调优信息');
    }
}

</script>

<style lang = 'less' scoped></style>
