import Image from "next/image";
import React from "react";
import { Button } from "../ui/button";
import { MapPin, UserRound } from "lucide-react";
import Logo from "@/public/logo.svg";
import Car from "@/types/car";
  
type SideBarProps = {
    selectedCar: Car | null;
    onSelectCar: (car: Car) => void;
    cars: Car[];
    dashboradPage: string;
    setDashboardPage: (page: string) => void;
}; 

export default function SideBar({cars, selectedCar, onSelectCar, dashboradPage, setDashboardPage }: SideBarProps) {
    return (
        <div className="h-full bg-[#131413] text-white w-[15%] py-2">
            <div className="divide-y w-full">
                <div className=" flex flex-col items-center py-8">
                    <Image width={95} src={Logo} alt="logo" />
                </div>
                <div className="py-8 px-6 flex flex-col items-start gap-y-6">
                    <button onClick={() => setDashboardPage("stations")}>
                        <div className="flex gap-3 items-center">
                            <MapPin className={`${dashboradPage == "stations" ? `text-[#7FEA52]` : `text-neutral-400`}`} size={18} />
                            <p className={`text-sm ${dashboradPage == "stations" ? `text-[#7FEA52]` : `text-neutral-400`}`}>Stations</p>
                        </div>
                    </button>

                    {/* <button onClick={() => setDashboardPage("account")}>
                        <div className="flex gap-3 items-center ">
                            <UserRound className={`${dashboradPage == "account" ? `text-[#7FEA52]` : `text-neutral-400`}`} size={18} />
                            <p className={`text-sm ${dashboradPage == "account" ? `text-[#7FEA52]` : `text-neutral-400`}`}>Account</p>
                        </div>
                    </button> */}

                </div>
                <div className="px-4 py-6">
                    <div>
                        <p className="font-semibold">My cars</p>
                        {cars.map((car, index) => (
                           <div key={index}
                                className="py-4">
                                <button 
                                onClick={() => onSelectCar(car)}
                                className={`bg-black w-[100%] h-32 rounded-md divide-y ${selectedCar?.name === car.name ? 'border-2 border-[#7FEA52]' : 'border-2 border-black'}`}>
                                    <div className="flex justify-around p-4">
                                        <Image src={car.image} alt="nissan" width={90} />
                                        <div className="flex flex-col">
                                            <p className="text-white text-xs">{car.name}</p>
                                            <p className="text-neutral-500 text-xs">{car.model}</p>
                                        </div>
                                    </div>
                                    <div className="flex py-2 justify-around">
                                        <div>
                                            <p className="text-neutral-500 text-xs">Battery</p>
                                            <p className="text-white text-xs">{car.battery}%</p>
                                        </div>
                                        <div>
                                            <p className="text-neutral-500 text-xs">Range</p>
                                            <p className="text-white text-xs">{car.range} Km</p>
                                        </div>
                                    </div>
                                </button>
                            </div> 
                        ))}
                    </div>
                </div>

            </div>
            {/* <div className="px-4 py-4">
                <div className="border border-black	bg-gradient-to-b from-black to-[#071609] w-[100%] h-32 rounded-md flex flex-col justify-center items-center px-4">
                    <p className="text-sm font-semibold">Sell Your Energy</p>
                    <p className="text-xs mb-4">Earn up to $250/week</p>
                    <Button>
                        Learn More
                    </Button>
                </div>
            </div> */}
        </div>
    );
}
