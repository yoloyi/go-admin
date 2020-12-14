import http from "@/utils/http.js";

import * as env from "@/utils/env.js";

export function userLogin(username, password) {
    return http.post(env.API_URL + "/auth/login", {
        username: username,
        password: password,
    })
}
