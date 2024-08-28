import { useState, useEffect } from "react";
import useHexToString from './useHexToString';
import axios from 'axios';
import Station from "@/types/station";

interface Report {
    payload: Station[];
}

interface InspectStations {
    status: string;
    exception_payload: any;
    reports: Report[];
    processed_input_count: number;
}

const useFetchStations = () => {

    const [stations, setStations] = useState<Station[]>([]);
    const [isLoadingStations, setIsLoadingStations] = useState(true); 

    useEffect(() => {
        const fetchStations = async () => {
            try {
                const response = await axios.get<InspectStations>('https://devolt.fly.dev/inspect/station');
                let jsonString = useHexToString(response.data.reports[0].payload);
                jsonString = jsonString.replace(/\u0000/g, '');
                const stations: Station[] = JSON.parse(jsonString);
                setStations(stations);
            } catch (error) {
                console.error('Error fetching data:', error);
            } finally {
                setIsLoadingStations(false);
            }
        };

        fetchStations();
    }, []);

    return { stations, isLoadingStations };
};

export default useFetchStations;
