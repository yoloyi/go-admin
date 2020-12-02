"use strict";

var _vue = _interopRequireDefault(require("vue"));

var _elementUi = _interopRequireDefault(require("element-ui"));

var _zhCN = _interopRequireDefault(require("element-ui/lib/locale/lang/zh-CN"));

require("../element-variables.scss");

function _interopRequireDefault(obj) { return obj && obj.__esModule ? obj : { "default": obj }; }

_vue["default"].use(_elementUi["default"], {
  locale: _zhCN["default"]
});