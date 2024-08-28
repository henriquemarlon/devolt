import { useState, useEffect } from "react";
import Station from "@/types/station";
import useFetchAddress from "@/hooks/useFetchAddress";
import useCalDistance from "@/hooks/useCalDistance";

const useStationManager = (stations: Station[], userLocation: string | [number, number]) => {

    const [stationsWithAddress, setStationsWithAddress] = useState<Station[]>([]);
    const [isLoadingAddresses, setIsLoadingAddresses] = useState(true);
    const [sortedStations, setSortedStations] = useState<Station[]>([]);

    useEffect(() => {
        const fetchAddresses = async () => {
            if (stations.length > 0) {
                try {
                    const updatedStations = await Promise.all(
                        stations.map(async (station) => {
                            const completeAddress = await useFetchAddress(station.latitude, station.longitude);
                            return { ...station, address: completeAddress };
                        })
                    );
                    setStationsWithAddress(updatedStations);
                } catch (error) {
                    console.error('Error fetching addresses:', error);
                } finally {
                    setIsLoadingAddresses(false);
                }
            }
        };

        fetchAddresses();
    }, [stations]);

    useEffect(() => {
        const stationsWithDistance = stationsWithAddress.map(station => {

            const distance = useCalDistance(
                Number(userLocation[0]),
                Number(userLocation[1]),
                station.latitude,
                station.longitude
            );

            return { ...station, distance };
        });

        const sortedStations = stationsWithDistance.sort((a, b) => a.distance - b.distance);
        setSortedStations(sortedStations);
    }, [userLocation, stationsWithAddress]);

    return { sortedStations, isLoadingAddresses };
}

export default useStationManager;
