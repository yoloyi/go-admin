"use strict";

Object.defineProperty(exports, "__esModule", {
  value: true
});
exports["default"] = void 0;

var _axios = _interopRequireDefault(require("axios"));

function _interopRequireDefault(obj) { return obj && obj.__esModule ? obj : { "default": obj }; }

var http = _axios["default"].create({
  baseURL: process.env.VUE_APP_API_URL,
  timeout: process.env.VUE_APP_API_TIMEOUT
});

http.interceptors.response.use(function (response) {
  return response;
}, function (error) {
  if (error.response.status === 401) {
    console.log(123);
  } // 对响应错误做点什么


  return Promise.reject(error);
});
var _default = http;
exports["default"] = _default;