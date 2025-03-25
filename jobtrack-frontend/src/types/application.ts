export enum Status {
    Sended = "Sended",
    Interview = "Interview",
    Reject = "Rejected",
    Accepted = "Accepted",
}

export interface Application {
    id: number;
    company: string;
    description: string;
    created_at: Date;
    status: Status;
}