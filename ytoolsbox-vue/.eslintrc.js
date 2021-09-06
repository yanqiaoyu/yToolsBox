module.exports = {
    root: true,
    parserOptions: {
      parser: 'babel-eslint',
      sourceType: 'module'
    },
    env: {
      browser: true,
      node: true,
      es6: true,
    },
    extends: ['plugin:vue/recommended', 'eslint:recommended'],
  
    // add your custom rules here
    //it is base on https://github.com/vuejs/eslint-config-vue
    rules: {
      "space-before-function-paren": 0,
      "vue/html-self-closing": ["error", {
        "html": {
          "void": "any",
          "normal": "never",
          "component": "any"
        },
        "svg": "any",
        "math": "any"
      }],
      "vue/max-attributes-per-line": ["warn", {
        "singleline": 10,    
        "multiline": {
          "max": 1,
          "allowFirstLine": false
        }
      }],
      "vue/singleline-html-element-content-newline": 0,
      "vue/multiline-html-element-content-newline":0,
      "vue/name-property-casing": ["error", "PascalCase"],
      "vue/attributes-order": 0
    }
  }