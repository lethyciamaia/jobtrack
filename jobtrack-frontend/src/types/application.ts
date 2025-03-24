export enum Status {
    Sended = "Sended",
    Interview = "Interview",
    Reject = "Reject",
    Accepted = "Accepted",
}

export interface Application {
    id: number;
    enterprise: string;
    description: string;
    createdAt: Date;
    status: Status;
}