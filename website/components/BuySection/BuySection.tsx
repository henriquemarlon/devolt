import { Button } from "../../components/ui/button";
import { MapPin, DollarSign, Zap, Info } from "lucide-react";
import { Slider } from "@/components/ui/slider"
import { Tooltip, TooltipContent, TooltipProvider, TooltipTrigger, } from "@/components/ui/tooltip"

type BuySectionProps = {
    energyPrice: string;
    percentage: number;
    carBattery: number;
    address: string;
    onBuyEnergy: (percentage: number) => void;
};

const BuySection: React.FC<BuySectionProps> = ({ address, energyPrice, percentage, carBattery, onBuyEnergy }) => (
    <div className="mt-12">
        <TooltipProvider>
            <div className="flex gap-2">
                <Zap className="text-[#7FEA52]" size={24} />
                <p className="text-xl">Buy Energy</p>
            </div>

            <Tooltip>
                <TooltipTrigger>
                    <div className="flex gap-2 mt-4">
                        <MapPin className="text-[#7FEA52]" size={18} />
                        <p className="text-sm pb-4">{address.length > 50
                            ? `${address.slice(0, 47)}...`
                            : address}</p>
                    </div>
                </TooltipTrigger>
                <TooltipContent>
                    <p>{address}</p>
                </TooltipContent>
            </Tooltip>

            <div className="flex gap-2">
                <DollarSign className="text-[#7FEA52]" size={18} />
                <p className="text-sm pb-4">{energyPrice}{energyPrice != "Select a station" ? "$ / $VOLT" : ""}</p>
            </div>
            <div className="flex gap-2 items-center">
                <Tooltip>
                    <TooltipTrigger>
                        <div className="flex items-center gap-1">
                            <p>{carBattery}%</p>
                            <Info className="text-[#7FEA52]" size={16} />
                        </div>
                    </TooltipTrigger>
                    <TooltipContent>
                        <p>Current Car Battery</p>
                    </TooltipContent>
                </Tooltip>
                <Slider onValueChange={(i) => onBuyEnergy(i[0])} defaultValue={[0]} max={100 - carBattery} step={1} />
                <p>100%</p>
            </div>

            <div className="flex justify-between mt-12 items-end">
                <div className="flex flex-col gap-2">
                    <div className="flex gap-2">
                        <Tooltip>
                            <TooltipTrigger>
                                <div className="flex items-center gap-1">
                                    <p>Final battery</p>
                                    <Info className="text-[#7FEA52]" size={16} />
                                    <p>:</p>
                                </div>
                            </TooltipTrigger>
                            <TooltipContent>
                                <p>Final battery after purchase</p>
                            </TooltipContent>
                        </Tooltip>
                        <p>{percentage + carBattery}%</p>
                    </div>
                    <div className="flex gap-2">
                        <p className="font-bold text-[#7FEA52]">Total price:</p>
                        <p>{energyPrice != "Select a station" ? (((percentage * 80) / 100) * Number(energyPrice)).toFixed(2) : "0"}$</p>
                    </div>
                </div>
                <Button disabled={energyPrice == "Select a station" ? true : false} className={`${energyPrice == "Select a station" ? `cursor-not-allowed bg-gray-500` : ``}`}>
                    Buy energy
                </Button>
            </div>
        </TooltipProvider>
    </div>
);
export default BuySection;