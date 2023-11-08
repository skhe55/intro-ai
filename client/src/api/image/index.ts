import { DEFAULT_API_PATH } from "$constants/index";
import type { TImage, TImageDto, TStandartApiResponse } from "$api/types";

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
    
    uploadImage = async (id: string, file: File): Promise<TStandartApiResponse<string> | null> => {
        try {
            const formData = new FormData();
            formData.append("file", file);

            const response = await fetch(`${DEFAULT_API_PATH}/images/upload/${id}`, {
                method: "POST",
                headers: {
                    'Authorization': `Bearer ${window.localStorage.getItem('token') as string}`,
                },
                body: formData,
            });

            if(response.ok) {
                return response.json();
            } else {
                return null;
            }
        } catch(e) {
            console.error(`uploadImage: ${e}`);
            return null;
        }
    };

    getImages = async (projectId: string): Promise<TStandartApiResponse<TImage[]> | null> => {
        try {
            const response = await fetch(`${DEFAULT_API_PATH}/images/${projectId}`, {
                headers: {
                    'Authorization': `Bearer ${window.localStorage.getItem('token') as string}`
                },
            });
            if (response.ok) {
                return response.json();
            } else {
                return null;
            }
        } catch(e) {
            console.error(`getImages: ${e}`);
            return null;
        }
    };
}