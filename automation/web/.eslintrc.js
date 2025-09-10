/******************************************************************************
 * Copyright (c) KylinSoft Co., Ltd. 2025.All rights reserved.
 * Non open source, The copyright belongs to KylinSoft Co., Ltd.
 ******************************************************************************/
module.exports = {
  root: true,
  env: {
    browser: true,
  },
  // 继承的配置
  extends: [
    "eslint:recommended",
    "plugin:@typescript-eslint/recommended",
    "plugin:vue/vue3-essential",
  ],
  parser: 'vue-eslint-parser',
  // 解析器选项
  parserOptions: {
    ecmaVersion: 2022,
    parser: "@typescript-eslint/parser",
    sourceType: 'module',
  },
  plugins: ["@typescript-eslint", "vue"],
  // 文件匹配规则
  overrides: [
    {
      env: {
        node: true,
      },
      files: [".eslintrc.{js,cjs}"],
      parserOptions: {
        sourceType: "script",
      },
    },
    {
      files: ['**/*.{ts,mts,tsx,vue}'],
      rules: {
        '@typescript-eslint/no-explicit-any': 0,
        '@typescript-eslint/no-unused-vars': 1,
        'vue/no-unused-vars': 1,
        '@typescript-eslint/no-unused-expressions': 1,
        '@typescript-eslint/no-empty-function': 1,
        'no-inner-declarations': 1,
        'no-undef': 1,
        'vue/multi-word-component-names': 0,
      },
    },
  ],
};
