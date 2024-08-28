import { Fuel } from "lucide-react";
import Station from "@/types/station";


type StationCardProps = {
    station: Station;
    setSelectStation: (station: Station) => void;
    selectedStation: Station | null;
};

const StationCard: React.FC<StationCardProps> = ({ station, setSelectStation, selectedStation }) => (
    
    <button onClick={() => setSelectStation(station)}>
        <div className={`w-56 h-60 p-6 bg-[#080908] rounded-md border-2 border-[#161d15] flex-shrink-0 ${selectedStation == station ? `border-[#7FEA52]`: ``}`}>
            <Fuel className="text-[#7FEA52]" size={32} />
            <div className="flex gap-1 pt-2">
                <p className="text-xl">{station.distance?.toString().slice(0, 4)}</p>
                <p className="text-xl text-neutral-500">Km</p>
            </div>
            <p className="pt-2 w-[95%] whitespace-normal text-sm text-left">{station.address?.length && station.address.length > 36
                ? `${station.address.slice(0, 38)}...`
                : station.address}</p>

            <div className="flex justify-start mt-4">  
                <div className="text-start">
                    <p className="text-xs text-neutral-500">Price per credit</p>
                    <p>$ {station.price_per_credit} $ / $VOLT</p>
                </div>
            </div>
        </div>
    </button>
);
export default StationCard;
