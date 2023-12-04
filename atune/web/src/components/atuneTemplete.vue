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
                <el-input v-model="form.note" class="custom-input" type="textarea" :rows="6"></el-input>
            </el-form-item>
        </el-form>
        <el-form class="centered-buttons">
            <el-button type="primary" @click="onSubmit" class="custom-button">保存</el-button>
            <el-button class="custom-button" @click="onCancel">取消</el-button>
        </el-form>
    </div>
</template>

<script lang='ts' setup>
import { ElMessage } from 'element-plus';
import { ref, onUpdated, reactive, onMounted } from 'vue'
import { getAtuneInfo, saveTune } from '@/api/atune'

let props = defineProps({
    selectedNodeData: {
        type: String,
        default: ""
    }
})
const atuneName = ref(props.selectedNodeData)
const emit = defineEmits(['closeDialog', 'dataUpdated']);

const form = reactive({
    tuneName: "",
    workDir: "",
    prepare: "",
    tune: "",
    restore: "",
    note: ""
})
const fetchAtuneInfo = () => {
    atuneName.value = props.selectedNodeData
    if (atuneName.value) {
        getAtuneInfo({ name: atuneName.value })
            .then((res) => {
                if (res.data && res.data.code === 200) {
                    const data = res.data.data;
                    // 判断数据结构类型
                    const baseTuneData = data.BaseTune || {};

                    form.tuneName = baseTuneData.tuneName || data.tuneName || '';
                    form.workDir = baseTuneData.workDir || data.workDir || '';
                    form.prepare = baseTuneData.prepare || data.prepare || '';
                    form.tune = baseTuneData.tune || data.tune || '';
                    form.restore = baseTuneData.restore || data.restore || '';
                    form.note = data.note || '';
                } else {
                    console.log('获取调优信息时出错:', res.data.msg)
                }
            })
    } else {
        console.warn('atuneName 为空，无法获取调优信息');
    }
}
const onSubmit = () => {
    emit('closeDialog');
    saveTune(form).then(res => {
        if (res.data.code === 200) {
            ElMessage.success(res.data.msg)
            emit('dataUpdated');
        } else {
            ElMessage.error(res.data.msg)
        }
    }).catch(err => {
        ElMessage.error("数据传输失败，请检查", err)
    })
}

const onCancel = () => {
    emit('closeDialog');
};

onMounted(() => {
    fetchAtuneInfo();
})

onUpdated(() => {
    fetchAtuneInfo();
})
</script>

<style lang = 'less' scoped>
.custom-form {
    margin-left: 25px;
    margin-right: 20px;
    outline-width: 120px;

    .custom-input {
        white-space: pre-wrap;
        resize: none;
        text-align: left;
        vertical-align: top;
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
