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
        <el-input
          v-model="form.note"
          class="custom-input"
          type="textarea"
          :rows="6"
        ></el-input>
      </el-form-item>
    </el-form>
    <el-form class="centered-buttons">
      <my-button v-show="isEdit" @click="onSubmit">保存</my-button>
      <my-button @click="onCancel">取消</my-button>
    </el-form>
  </div>
</template>

<script lang="ts" setup>
import { ElMessage } from "element-plus";
import { reactive, watchEffect, ref } from "vue";
import { getAtuneInfo, saveTune, updateTune } from "@/api/atune";
import { Atune } from "@/types/atune";
import MyButton from "@/components/myButton.vue";

let props = defineProps({
  selectedNodeData: {
    type: String,
    default: "",
  },
  selectedEditRow: {
    type: Object as () => Atune,
    default: null,
  },
});

const isEdit = ref(false);
const form = reactive({
  id: 0,
  tuneName: "",
  workDir: "",
  prepare: "",
  tune: "",
  restore: "",
  note: "",
});
const emit = defineEmits(["closeDialog", "dataUpdated"]);

// 编辑
const handleEdit = () => {
  form.id = props.selectedEditRow.id;
  form.tuneName = props.selectedEditRow.tuneName;
  form.workDir = props.selectedEditRow.workDir;
  form.prepare = props.selectedEditRow.prepare;
  form.tune = props.selectedEditRow.tune;
  form.restore = props.selectedEditRow.restore;
  form.note = props.selectedEditRow.note;
};

const fetchAtuneInfo = () => {
  getAtuneInfo({ name: props.selectedNodeData }).then((res) => {
    if (res.data && res.data.code === 200) {
      const data = res.data.data;
      // 判断数据结构类型
      const baseTuneData = data.BaseTune || {};

      form.tuneName = baseTuneData.tuneName || data.tuneName || "";
      form.workDir = baseTuneData.workDir || data.workDir || "";
      form.prepare = baseTuneData.prepare || data.prepare || "";
      form.tune = baseTuneData.tune || data.tune || "";
      form.restore = baseTuneData.restore || data.restore || "";
      form.note = data.note || "";
    } else {
      console.log("获取调优信息时出错:", res.data.msg);
    }
  });
};

const saveTuneData = () => {
  saveTune(form)
    .then((res) => {
      if (res.data.code === 200) {
        ElMessage.success(res.data.msg);
        emit("dataUpdated");
      } else {
        ElMessage.error(res.data.msg);
      }
    })
    .catch((err) => {
      ElMessage.error("数据传输失败，请检查", err);
    });
};
const updateTuneData = () => {
  updateTune(form)
    .then((res) => {
      if (res.data.code === 200) {
        ElMessage.success(res.data.msg);
        emit("dataUpdated");
      } else {
        ElMessage.error(res.data.msg);
      }
    })
    .catch((err) => {
      ElMessage.error("数据传输失败，请检查", err);
    });
};

const onSubmit = () => {
  emit("closeDialog");
  if (form.id === 0) {
    saveTuneData();
  } else {
    updateTuneData();
  }
};

const onCancel = () => {
  emit("closeDialog");
};

watchEffect(() => {
  if (props.selectedNodeData !== "") {
    fetchAtuneInfo();
  }
});

watchEffect(() => {
  if (props.selectedEditRow) {
    handleEdit();
  }
});
</script>

<style lang="less" scoped>
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
</style>
