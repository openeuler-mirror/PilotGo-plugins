<template>
  <div class="cont">
    <div v-show="isChart" class='echart' ref="chartDom"></div>
    <span v-show="!isChart && (type === 'value')" class="text">{{ char_value }}</span>
    <chart-table class="table" v-show="!isChart && (type==='table')" :columnList="columnList" :tableData="tableData" />
  </div>
</template>
  
<script setup lang='ts' scoped>
import { inject, ref, onMounted, reactive, watch, markRaw, nextTick } from 'vue'
import { getPromeCurrent, getPromeRange } from '@/api/prometheus';
import ChartTable from './chartTable.vue';
import { filterProm, deepClone, handle_byte, startTime, endTime, line_opt, gauge_opt } from './index';
import { formatDate } from '@/utils/dateFormat';
import { useMacStore } from '@/store/mac';
let macIp = ref('');
const echarts: any = inject('echarts')
const chartDom = ref(null)
const option = ref({});
const myChart = ref<any>(null);
const char_value = ref('0.00');
const line_arr = ref([] as any[]); // 存放多请求的折线图数据集合
let line_option = reactive(deepClone(line_opt));
let gauge_option = reactive(deepClone(gauge_opt));

interface tableItem {
  prop: string,
  label: string
}
const columnList = ref([] as tableItem[])
const tableData = ref([] as any)
const props = defineProps({
  query: {
    type: Object,
    default: {},
    required: true
  },
})
const search_step = 15; // 查询区间数据的步长
const isChart = props.query.isChart;

const type = props.query.type;
const resize = () => {
  // 由于网格布局拖拽放大缩小图表不能自适应，这里设置一个定时器使得echart加载为一个异步过程
  setTimeout(() => {
    nextTick(() => {
      myChart.value.resize()
    })
  }, 0);
}
// 获取prometheus数据
const getPromeData = (item: any) => {
  if (item.range) {
    let proms = [] as any;
    item.sqls.forEach((sqlItem: any, index: number) => {
      // 1.使用new promise来进行异步处理,避免乱序
      proms.push(
        new Promise((resolve, reject) => {
          getPromeRange({ query: sqlItem.sql.replace(/{macIp}/g, macIp.value), start: sqlItem.start || '', end: sqlItem.end || endTime, step: sqlItem.step || search_step }).then(res => {
            return resolve(filterProm(res) && filterProm(res))
          }).catch(err => {
            return reject(err)
          })
        })
      )
    });
    // 2. 使用promise.all拿到所有的数据
    Promise.all(proms).then(res => {
      line_arr.value = res;
    })
  }
  else if (!item.range) {
    let proms = [] as any;
    item.sqls.forEach(async (sqlItem: any, index: number) => {
      proms.push(
        new Promise((resolve, reject) => {
          getPromeCurrent({ query: sqlItem.sql.replace(/{macIp}/g, macIp.value) }).then(res => {
            return resolve(filterProm(res) && filterProm(res))
          }).catch(err => {
            return reject(err)
          })
        })
      )

    });
    Promise.all(proms).then(res => filterCurrentData(item, res))
  }
}

// 过滤基础数据类型
const filterCurrentData = (item: any, result: any) => {
  switch (item.type) {
    case 'value':
      set_value_type(item, result[0]);
      break;
    case 'gauge':
      set_gauge_type(item, result[0]);
      break;
    case 'table':
      if (result.length != props.query.sqls.length) {
        return false;
      }
      set_table_type(item, result)
      break;
    default:
      break;
  }
}

// 处理折线图data数据
const handle_line_data = (values: any, target: string) => {
  let line_data = [] as any;
  values.forEach((valueItem: any) => {
    let time_text = formatDate(new Date(valueItem[0] * 1000), "YYYY-MM-DD HH:ii:ss")
    let item_value = '';
    switch (target) {
      case 'byte_series':
        item_value = handle_byte(valueItem[1], 2);
        break;
      case 'percent_series':
        item_value = (parseFloat(valueItem[1]) * 100).toFixed(2);
        break;
      case 'speed_series':
        item_value = parseFloat(valueItem[1]).toFixed(2);
        break;
      default:
        item_value = parseFloat(valueItem[1]).toFixed(2);
        break;
    }
    line_data.push({
      time: valueItem[0] * 1000,
      value: [time_text, item_value]
    })
  })
  return line_data;
}

// 设置折线类型的数据
const set_line_type = (item: any, result: any) => {
  line_option.yAxis.axisLabel.formatter = '{value}' + item.unit;
  let series = {
    name: '', type: 'line', smooth: false, showSymbol: false,
    z: 1, zlevel: 1, lineStyle: { width: 1 }, areaStyle: { opacity: 0.1, }, data: [],
  } as any;
  line_arr.value.forEach((line: any, lineIndex: number) => {
    let legendName = deepClone(item.sqls[lineIndex].series_name);
    if (line instanceof Array) {
      // 如果是数组
      line.forEach(lineItem => {
        series.name = lineItem.metric && deepClone(lineItem.metric.device) || '';
        series.data = deepClone(handle_line_data(lineItem.values, item.target));
        line_option.series.push(deepClone(series))
      })
    } else {
      series.name = legendName;
      series.data = deepClone(handle_line_data(line.values, item.target));
      line_option.series.push(deepClone(series))
    }
  })
  option.value = line_option;
}

// 设置数值类型的数据
const set_value_type = (item: any, result: any) => {
  switch (item.target) {
    case 'value_series':
      // 数值系列
      char_value.value = result && result.value ? parseFloat(result.value[1]).toFixed(item.float) + item.unit : '0';
      break;
    case 'byte_series':
      // 字节系列
      char_value.value = result && result.value ? handle_byte(result.value[1], item.float) + item.unit : '0';
      break;

    default:
      break;
  }
}

// 设置仪表盘类型的数据
const set_gauge_type = (item: any, result: any) => {
  gauge_option.series[0].min = item.min || 0;
  gauge_option.series[0].max = item.max || 100;
  gauge_option.series[0].axisLine.lineStyle.color = item.color;
  gauge_option.series[0].detail.formatter = '{value}' + item.unit;
  switch (item.target) {
    case 'percent_series':
      // 百分比系列
      gauge_option.series[0].data[0].value = result && result.value ? parseFloat(result.value[1]).toFixed(item.float) : 0;
      break;
    case 'num_series':
      // 数值系列
      let num = result && result.value ? parseFloat(result.value[1]) / 1000 : 0;
      if (num <= 1) {
        gauge_option.series[0].data[0].value = num * 1000;//(parseInt(result.value[1]) / 1000).toFixed(item.float);
      } else {
        gauge_option.series[0].detail.formatter = '{value}K'
        gauge_option.series[0].data[0].value = num.toFixed(item.float);
      }
      break;
    default:
      break;
  }
  option.value = gauge_option;
}

// 设置表格类型的数据
const set_table_type = (item: any, result: any) => {
  tableData.value = [];
  columnList.value = item.sqls[0].columnList;
  let colList = [] as any[];
  item.sqls.forEach((item: any) => {
    colList.push(...item.columnList)
  })
  columnList.value = colList;
  let tableData1 = {} as any;
  item.sqls.forEach((sqlItem: any, index: number) => {
    let cols = deepClone(sqlItem.columnValue);
    let tableItem = {} as any;
    result[index].forEach((item: any, resIndex: number) => {
      cols.forEach(async (vItem: any) => {
        switch (vItem.type) {
          case 'byte':
            tableItem[vItem.prop] = handle_byte(eval(vItem.value), 2)
            break;
          case 'percent':
            tableItem[vItem.prop] = (parseFloat(eval(vItem.value)) * 100).toFixed(2) + '%'
            break;
          case 'float':
            tableItem[vItem.prop] = parseFloat(eval(vItem.value)).toFixed(2)
            break;
          default:
            tableItem[vItem.prop] = eval(vItem.value)
            break;
        }
      })
      tableData1[`res${resIndex}.${index}`] = deepClone(tableItem);
    })
  })
  result[0].forEach((key: any, index: number) => {
    let rowData = {} as any;
    for (let i = 0; i < item.sqls.length; i++) {
      let resString = 'res' + index + '.' + i;
      rowData = Object.assign(rowData, tableData1[resString])
    }
    tableData.value.push(deepClone(rowData))
  })

}

/* setInterval(() => {
  getPromeData(props.query)
}, 5000); */

onMounted(() => {
  myChart.value = markRaw(echarts.init(chartDom.value))
  if (props.query.isChart) {
    myChart.value.setOption(option.value, true)
  }
  setTimeout(() => {
    macIp.value = useMacStore().macIp;
    getPromeData(props.query);
  }, 8)
  window.addEventListener('resize', resize)
})

watch(() => option.value, (new_option) => {
  if (myChart.value.getOption()) {
    myChart.value.dispose();
    myChart.value = markRaw(echarts.init(chartDom.value))
  }
  myChart.value.setOption(new_option, true)
}, {
  deep: true
})

watch(() => line_arr.value, (new_line_arr) => {
  if (new_line_arr.length != props.query.sqls.length) {
    return false;
  }
  set_line_type(props.query, new_line_arr);
}, {
  deep: true
})
defineExpose({
  resize
})
</script>
  
<style lang='scss' scoped>
.cont {
  width: 100%;
  height: 100%;

  .echart,
  .text,
  .table {
    width: 100%;
    height: 100%;
  }

  .text {
    display: flex;
    justify-content: center;
    align-items: center;
    font-weight: bold;
    font-size: 20px;
    color: #67e0e3;
    user-select: none;
  }
}
</style>