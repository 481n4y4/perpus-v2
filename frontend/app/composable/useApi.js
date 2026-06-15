export const useApi = () => {
    const apiBase = useRuntimeConfig().public.apiBase;
    const token = useCookie("auth_token");

    return $fetch.create({
        baseURL: apiBase,
        onRequest({options}) {
            options.headers = {
                Accept: "application/json",
                ...(options.headers || {}),
                ...(token.value ? {Authorization: `Bearer ${token.value}`} : {})
            }
        }
    })
}