<template>
  <h1 class="h1">单机拓扑图演示页面</h1>
  <div id="topo-container" class="container"></div>

  <el-drawer class="drawer" v-model="drawer" :title="title" :direction="direction" :before-close="handleClose">
    <span>Hi, there!</span>
  </el-drawer>
</template>

<script setup lang="ts">
import G6 from '@antv/g6';
import { ref, onMounted } from "vue";


let drawer = ref(false)
const direction = 'rtl'
const title = ref("")

function handleClose() {
  drawer.value = false
}

onMounted(() => {
  initGraph()

})

function initGraph() {
  const data = {
    "id": "Modeling Methods",
    "children": [
      {
        "id": "Classification",
        "children": [
          {
            "id": "Logistic regression"
          },
          {
            "id": "Linear discriminant analysis"
          },
          {
            "id": "Rules"
          },
          {
            "id": "Decision trees"
          },
          {
            "id": "Naive Bayes"
          },
          {
            "id": "K nearest neighbor"
          },
          {
            "id": "Probabilistic neural network"
          },
          {
            "id": "Support vector machine"
          }
        ]
      },
      {
        "id": "Consensus",
        "children": [
          {
            "id": "Models diversity",
            "children": [
              {
                "id": "Different initializations"
              },
              {
                "id": "Different parameter choices"
              },
              {
                "id": "Different architectures"
              },
              {
                "id": "Different modeling methods"
              },
              {
                "id": "Different training sets"
              },
              {
                "id": "Different feature sets"
              }
            ]
          },
          {
            "id": "Methods",
            "children": [
              {
                "id": "Classifier selection"
              },
              {
                "id": "Classifier fusion"
              }
            ]
          },
          {
            "id": "Common",
            "children": [
              {
                "id": "Bagging"
              },
              {
                "id": "Boosting"
              },
              {
                "id": "AdaBoost"
              }
            ]
          }
        ]
      },
      {
        "id": "Regression",
        "children": [
          {
            "id": "Multiple linear regression"
          },
          {
            "id": "Partial least squares"
          },
          {
            "id": "Multi-layer feedforward neural network"
          },
          {
            "id": "General regression neural network"
          },
          {
            "id": "Support vector regression"
          }
        ]
      }
    ]
  }

  let graph = new G6.TreeGraph({
    container: "topo-container",
    width: document.getElementById("topo-container")!.clientWidth,
    height: document.getElementById("topo-container")!.clientHeight,
    modes: {
      default: ['drag-canvas', 'zoom-canvas', "click-select",
       ],
    },
    layout: {
      type: 'dendrogram',
      direction: 'LR', // H / V / LR / RL / TB / BT
      nodeSep: 30,
      rankSep: 100,
    },
  });
  graph.node(function (node) {
      return {
        label: node.id,
        labelCfg: {
          position: node.children && node.children.length > 0 ? 'left' : 'right',
          offset: 5,
        },
      };
    });
  graph.on("nodeselectchange", (e) => {
    if (e.select) {
      let node_id = e.target._cfg!.id
      console.log("click node:", node_id);

      title.value = "I am " + node_id;
      drawer.value = drawer.value ? false : true;
    } else {
      console.log("node unselected")
    }
    return false
  });
  graph.data(data);
  graph.render();
  graph.fitCenter();

  window.onresize = () => {
    graph.changeSize(
      document.getElementById("topo-container")!.clientWidth,
      document.getElementById("topo-container")!.clientHeight)
    graph.fitCenter()
  }
}

</script>

<style scoped>
.h1 {
  width: 100%;
  text-align: center;
}

.container {
  width: 100%;
  height: 100%;
  background-color: white;
}

.drawer {
  height: 100%;
}
</style>
