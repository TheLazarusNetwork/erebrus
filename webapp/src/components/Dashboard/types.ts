// types.ts
export interface Client {
    UUID: string;
    Name: string;
    Tags: string[];
    Email: string;
    Enable: boolean;
    PresharedKey: string;
    AllowedIPs: string[];
    Address: string[];
    PrivateKey: string;
    PublicKey: string;
    CreatedBy: string;
    Created: number;
    Updated: number;
  }
  