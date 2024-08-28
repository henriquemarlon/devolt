import React, { useState } from "react";
import { MapPin } from "lucide-react";
import BuySection from "@/components/BuySection/BuySection";
import StationCard from "@/components/StationCard/StationCard";
import VehicleStats from "@/components/VehicleStats/VehicleStats";
import Car from "@/types/car";
import useWindowSize from "@/hooks/useWindowSize";
import useMap from "@/hooks/useMap";
import useUserLocation from "@/hooks/useUserLocation";
import Station from "@/types/station";
import BeatLoader from "react-spinners/BeatLoader";
import useFetchStations from "@/hooks/useFetchStations";
import useStationManager from "@/hooks/useStationManager";

type DashboardStationProps = {
    selectedCar: Car;
};

export default function DashboardStation({ selectedCar: car }: DashboardStationProps) {

    const Map = useMap();
    const windowSize = useWindowSize();
    const [userLocation, city] = useUserLocation();

    const [percentage, setPercentage] = useState(0);  
    const [selectedStation, setSelectedStation] = useState<Station | null>(null);
    
    const mapSize = windowSize.height <= 800 ? "300px" : "400px";

    const { stations, isLoadingStations } = useFetchStations();
    const { sortedStations, isLoadingAddresses }: { sortedStations: Station[], isLoadingAddresses: boolean } = useStationManager(stations, userLocation);


    return (
        <div className="w-[80%] m-8">
            <div className="w-[100%] flex justify-between pb-6">
                <p className="text-2xl font-semibold">Stations</p>
                <w3m-button />
            </div>
            <div className="flex gap-8">
                <div className="w-[60%]">
                    <div className="w-full max-w-full bg-[#080908] rounded-md border-2 border-[#161d15] h-[60%] 2xl:h-[65%] p-6">
                        <div className="flex gap-2">
                            <MapPin className="text-[#7FEA52]" size={24} />
                            <p className="text-xl pb-6">{city !== "" ? city : "Location"}</p>
                        </div>
                        <Map
                            center={userLocation}
                            stations={sortedStations}
                            height={mapSize}
                            width="100%"
                            roundedBottomCorners
                            roundedTopCorners
                            userLocation={userLocation}
                        />
                    </div>
                    <div className="w-full overflow-x-scroll whitespace-nowrap mt-6 h-[40%] custom-scrollbar">
                        <div className="flex flex-nowrap space-x-4">
                            {isLoadingStations || isLoadingAddresses ? (
                               <BeatLoader color="#7FEA52" />
                            ) : (
                                sortedStations.map((station, index) => (
                                    <StationCard
                                        key={index}
                                        station={station}
                                        setSelectStation={setSelectedStation}
                                        selectedStation={selectedStation}
                                    />
                                ))
                            )}
                        </div>
                    </div>
                </div>
                <div className="w-[40%] bg-[#080908] rounded-md border-2 border-[#161d15] p-6">
                    <VehicleStats car={car} />
                    <BuySection
                        energyPrice={selectedStation?.price_per_credit || "Select a station"}
                        percentage={percentage}
                        carBattery={car.battery}
                        address={selectedStation?.address || "Select a station"}
                        onBuyEnergy={(i) => setPercentage(i)}
                    />
                </div>
            </div>
        </div>
    );
}
