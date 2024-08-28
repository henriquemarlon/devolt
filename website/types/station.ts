export default interface Station {
    id: number;
    consumption: string;
    owner: string;
    price_per_credit: string;
    state: string;
    latitude: number;
    longitude: number;
    created_at: number;
    updated_at: number;
    distance?: number | undefined;
    address?: string | undefined | null;
  }