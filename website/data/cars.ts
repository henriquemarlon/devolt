import Car from "@/types/car";
import NissanCar from "@/public/nissan-car.svg";
import TeslaCar from "@/public/tesla-car.svg";


const cars: Car[] = [
    {
        name: "Nissan",
        model: "Leaf",
        image: NissanCar,
        battery: 78,
        range: 230,
        capacity: 90,
    },
    {
        name: "Tesla",
        model: "Model X",
        image: TeslaCar,
        battery: 89,
        range: 300,
        capacity: 80,
    },
];

export default cars;