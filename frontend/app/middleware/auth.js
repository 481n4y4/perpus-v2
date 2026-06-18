export default defineNuxtRouteMiddleware(async (to, from) => {
    const token = useCookie("auth_token")
    if (!token.value && to.path === '/auth/login') {
        return
    }
    if (!token.value) {
        return navigateTo("/auth/login")
    }
})