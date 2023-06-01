
export let startTime = (new Date() as any) / 1000 - 60 * 60 * 2;
export let endTime = (new Date() as any) / 1000;

// 过滤prometheus接口返回的数据
export const filterProm = (res: any) => {
  if (res.data.status === 'success' && res.data.data.result.length > 0) {
    return res.data.data.result.length === 1 ? res.data.data.result[0] : res.data.data.result;
  } else {
    return [];
  }
}
// 深拷贝数组
export const deepClone = (res: any) => {
  return JSON.parse(JSON.stringify(res))
}

// 处理字节类型数据
export const handle_byte = (value: string, float: number, unit: string) => {
  let value2unit: string = '0.00';
  switch (unit) {
    case 'GB':
      value2unit = (parseFloat(value) / 1024 / 1024 / 1024).toFixed(float);
      break;
    case 'KB':
      value2unit = (parseFloat(value) / 1024).toFixed(float);
      break;

    default:
      // MB
      value2unit = (parseFloat(value) / 1024 / 1024).toFixed(float);
      break;
  }
  return value2unit;
}

// echart颜色表
let e_colors = ['#4980c9', '#67e0e3', '#ef6874', '#4ab92e', '#E6A23C', '#74a465', '#e24d42', '#ba43a9'];

// 柱状图
export const bar_opt = {
  color: e_colors,
  xAxis: {
    type: 'category',
    data: ['Mon', 'Tue', 'Wed', 'Thu', 'Fri', 'Sat', 'Sun']
  },
  yAxis: {
    type: 'value'
  },
  series: [
    {
      data: [120, 200, 150, 80, 70, 110, 130],
      type: 'bar',
      showBackground: true,
      backgroundStyle: {
        color: 'rgba(180, 180, 180, 0.2)'
      }
    }
  ]
}

// 折线图
export const line_opt = {
  color: e_colors,
  tooltip: {
    trigger: 'axis',
    position: [10, 60],
    formatter: function (params: any) {
      params = params[0];
      return (
        params.value[0] + ' ' +
        params.value[1]
      );
    },
    axisPointer: {
      animation: false
    }
  },
  legend: {
    show: true,
    orient: 'vertical',
    right: 10,
    top: 20,
    textStyle: {
      color: '#888'
    },
    icon: 'roundRect',
  },
  grid: {
    top: '4%',
    left: '2%',
    right: '20%',
    bottom: '2%',
    containLabel: true
  },
  xAxis: {
    type: 'time',
    splitLine: {
      show: false
    }
  },
  yAxis: {
    type: 'value',
    splitLine: {
      show: true,
      lineStyle: {
        color: ['#333']
      }
    },
    min: 0,
    boundaryGap: [0, '100%'],
    axisLabel: {
      formatter: '{value}'
    }
  },
  series: [{
    name: '',
    type: 'line',
    smooth: false,
    showSymbol: false,
    z: 1,
    zlevel: 1,
    lineStyle: {
      width: 1
    },
    areaStyle: {
      opacity: 0.1,
    },
    data: [],
  }]
}
// 仪表盘
export const gauge_opt = {
  series: [
    {
      type: 'gauge',
      center: ['50%', '60%'],
      radius: '110%',
      startAngle: 200,
      endAngle: -20,
      min: 0,
      max: 100,
      tooltip: {

      },
      progress: {
        show: false,
      },
      pointer: {
        show: true,
        itemStyle: {
          color: 'inherit'
        }
      },
      axisLine: {
        lineStyle: {
          width: 6,
          color: [
            [0.3, '#67e0e3'],
            [0.7, '#E6A23C'],
            [1, '#fd666d']
          ]
        }
      },
      axisTick: {
        show: false,
      },
      splitLine: {
        show: false
      },
      axisLabel: {
        show: false
      },
      anchor: {
        show: false
      },
      title: {
        show: false
      },
      detail: {
        valueAnimation: true,
        offsetCenter: [0, '50%'],
        fontSize: 20,
        fontWeight: 'bolder',
        formatter: '{value}',
        color: 'inherit'
      },
      data: [
        {
          value: 64.40
        }
      ]
    },
  ]
};