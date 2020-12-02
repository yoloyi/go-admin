import axios from 'axios';

const http = axios.create(
	{
		baseURL: process.env.VUE_APP_API_URL,
		timeout: process.env.VUE_APP_API_TIMEOUT
	}
)

http.interceptors.response.use(function (response) {
	return response;
}, function (error) {
	if (error.response.status === 401) {
		console.log(123)
	}
	// 对响应错误做点什么
	return Promise.reject(error);
});

export default http

