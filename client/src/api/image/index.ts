import { DEFAULT_API_PATH } from "$constants/index";
import type { TImageDto, TStandartApiResponse } from "$api/types";

export class ImageApi {
    createImage = async (payload: TImageDto): Promise<TStandartApiResponse<string> | null> => {
        try {
            const response = await fetch(`${DEFAULT_API_PATH}/images/create`, {
                method: "POST",
                headers: {
                    'Authorization': `Bearer ${window.localStorage.getItem('token') as string}`,
                    'Content-Type': 'application/json'
                },
                body: JSON.stringify(payload),
            });
            if(response.ok) {
                return response.json();
            } else {
                return null;
            }
        } catch(e) {
            console.error(`createImage: ${e}`);
            return null;
        }
    };
    
    uploadImage = async (id: string) => {
        try {
            const response = await fetch(`${DEFAULT_API_PATH}/images/upload/${id}`, {
                method: "POST",
                headers: {
                    'Authorization': `Bearer ${window.localStorage.getItem('token') as string}`,
                    'Content-Type': 'multipart/form-data'
                },
            })
        } catch(e) {
            console.error(`uploadImage: ${e}`);
            return null;
        }
    }
}