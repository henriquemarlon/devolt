import Payload from "@/types/station";


export default interface MapProps {
    stations: Payload[];
    width?: string;
    height?: string;
    center: any;
    userLocation: any;
    buttonText?: string;
    roundedTopCorners: boolean;
    roundedBottomCorners: boolean;
    setSelectedStation?: (station: any) => void;
    hidden?: boolean;
  }