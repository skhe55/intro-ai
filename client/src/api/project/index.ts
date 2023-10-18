import { HEADERS_WITH_BEARER_TOKEN, DEFAULT_API_PATH } from "$constants/index";
import type { TProject, TProjectDto, TStandartApiResponse } from "$api/types";

export class ProjectApi {
    getProjects = async (): Promise<TStandartApiResponse<TProject[]> | null> => {
        try {
            const response = await fetch(`${DEFAULT_API_PATH}/projects`, HEADERS_WITH_BEARER_TOKEN(window.localStorage.getItem('token') as string));
            if(response.ok) {
                return response.json();
            } else {
                return null;
            }
        } catch(e) {
            console.error(`getProjects: ${e}`);
            return null;
        }
    };

    createProject = async (payload: TProjectDto): Promise<TStandartApiResponse<string> | null> => {
        try {
            const response = await fetch(`${DEFAULT_API_PATH}/projects/create`, {
                method: "POST",
                ...HEADERS_WITH_BEARER_TOKEN(window.localStorage.getItem('token') as string, true),
                body: JSON.stringify(payload),
            });
            if(response.ok) {
                return response.json();
            } else {
                return null;
            }
        } catch(e) {
            console.error(`createProject: ${e}`);
            return null;
        }
    };

    deleteProject = async (id: string): Promise<TStandartApiResponse<string> | null> => {
        try {
            const response = await fetch(`${DEFAULT_API_PATH}/projects/delete/${id}`, {
                method: "DELETE",
            });
            if(response.ok) {
                return response.json();
            } else {
                return null;
            }
        } catch(e) {
            console.error(`deleteProject: ${e}`);
            return null;
        }
    };
}