import { CarFront } from "lucide-react";
import Image from "next/image";
import Car from "@/types/car";

interface VehicleStatsProps {
    car: Car;
}

const VehicleStats: React.FC<VehicleStatsProps> =  ({car}) => {
    return <div>
        <div className="flex gap-2">
            <CarFront className="text-[#7FEA52]" size={24} />
            <p className="text-xl">Vehicle Stats</p>
        </div>
        <p className="pb-4 text-neutral-500">{car.model}</p>
        <div className="flex justify-center mt-12 h-32 2xl:h-48">
            <Image width={350} src={car.image} alt="logo" />
        </div>
        <div className="flex justify-around mt-8">
            <div className="w-20">
                <p className="text-neutral-400">EV</p>
                <p className="text-lg">{car.name}</p>
            </div>
            <div className="w-20">
                <p className="text-neutral-400">Battery</p>
                <p className="text-lg">{car.battery}%</p>
            </div>
            <div className="w-20">
                <p className="text-neutral-400">Range</p>
                <p className="text-lg">{car.range} Km</p>
            </div>
            <div className="w-20">
                <p className="text-neutral-400">Capacity</p>
                <p className="text-lg">{car.capacity} kWh</p>
            </div>
        </div>
    </div>;
}

export default VehicleStats;