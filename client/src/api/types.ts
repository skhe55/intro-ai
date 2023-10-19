export type TStandartApiResponse<T> = {
    Message: string;
    Status: string;
    Result: T;
}

export type TSignInPayload = {
    login: string;
    password: string;
};

export type TSignInUpResponse = {
    username: string;
    token: string;
    expires_at: string;
};

export type TSignUpPayload = TSignInPayload & { username?: string };

// Projects

export type TProject = {
    id: string;
    name: string;
    images: TImage[] | null;
};

export type TProjectDto = {
    name: string;
};

// Images

export type TImage = {
    id: string;
    filename: string;
    pathToImage: string;
    projectId: string;
    coordinates: number[][] | null;
    created_at: Date;
};

export type TImageDto = {
    filename: string;
    projectId: string;
    coordinates: number[][];
};