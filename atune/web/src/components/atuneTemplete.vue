<template>
  <div>
    <el-form :model="form" class="custom-form" :disabled="!isTune">
      <el-form-item label="调优对象" v-show="isTune">
        <el-select
          v-model="form.tuneName"
          placeholder="请选择调优模板"
          @change="fetchAtuneInfo(form.tuneName)"
        >
          <el-option
            v-for="item in allTune"
            :key="item"
            :label="item"
            :value="item"
          />
        </el-select>
      </el-form-item>
      <el-form-item label="自定义名">
        <el-input v-model="form.custom_name"></el-input>
      </el-form-item>
      <el-form-item label="概述介绍">
        <el-input v-model="form.description"></el-input>
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
      <my-button v-show="isTune" @click="onSubmit">保存</my-button>
      <my-button @click="onCancel">取消</my-button>
    </el-form>
  </div>
</template>

<script lang="ts" setup>
import { ElMessage } from "element-plus";
import { reactive, watchEffect, ref } from "vue";
import {
  getAtuneInfo,
  saveTune,
  updateTune,
  getAtuneAllName,
} from "@/api/atune";
import { Atune, Task } from "@/types/atune";
import MyButton from "@/components/myButton.vue";

let props = defineProps({
  isTune: {
    type: Boolean,
    default: false,
    required: true,
  },
  selectedEditRow: {
    type: Object as () => Atune,
    default: null,
  },
});
const allTune = ref([""]);
const form = reactive({
  id: 0,
  tuneName: "",
  custom_name: "",
  description: "",
  workDir: "",
  prepare: "",
  tune: "",
  restore: "",
  note: "",
});
const emit = defineEmits(["closeDialog"]);
// 获取所有的调优模板
const getAllTune = () => {
  getAtuneAllName().then((res) => {
    if (res.data.code === 200) {
      allTune.value = res.data.data;
    }
  });
};

// 编辑
const handleEdit = () => {
  form.id = props.selectedEditRow.id;
  form.tuneName = props.selectedEditRow.tuneName;
  form.custom_name = props.selectedEditRow.custom_name;
  form.workDir = props.selectedEditRow.workDir;
  form.prepare = props.selectedEditRow.prepare;
  form.tune = props.selectedEditRow.tune;
  form.restore = props.selectedEditRow.restore;
  form.note = props.selectedEditRow.note;
  form.description = props.selectedEditRow.description;
};

const fetchAtuneInfo = (tuneName: string) => {
  getAtuneInfo({ name: tuneName }).then((res) => {
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
  if (props.isTune) {
    getAllTune();
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
.centered-buttons {
  text-align: center;
}
</style>
