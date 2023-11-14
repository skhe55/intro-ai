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
    name: string;
    path_to_image: string;
    projectId: string;
    created_at: Date;
};

export type TImageDto = {
    name: string;
    projectId: string;
};

export type TImageDeleteDto = {
    project_id: string;
    path_to_image: string;
};

export type TLabelDto = {
    image_id: string;
    name: string;
};

export type TLabel = {
    id: string;
    name: string;
    image_id: string;
    created_at: Date;
};