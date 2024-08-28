import dynamic from "next/dynamic";
import React, { useMemo } from "react";
import BeatLoader from "react-spinners/BounceLoader";


export default function useMap() {
    const Map = useMemo(
        () =>
            dynamic(() => import("@/components/Map/Map"), {
                loading: () => <div className="h-56 w-auto flex justify-center items-center"><BeatLoader color="#7FEA52" /></div>,
                ssr: false,
            }),
        []
    );

    return Map;
}