import { API } from '@/service/api/api';

export const login = async (credentials: { email: string; password: string; remember: boolean }) => {
    try {
        // await WEB.get(`/sanctum/csrf-cookie`);
        return API.post('/login', credentials);
    } catch (error) {
        throw error instanceof Error ? error : new Error(String(error));
    }
};
