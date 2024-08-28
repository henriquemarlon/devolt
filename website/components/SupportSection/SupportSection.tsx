import React, { useState } from "react";
import { Card } from "../ui/card";
import { Button } from "../ui/button";
import { SubmitHandler, useForm } from "react-hook-form";
import axios from "axios";
import { toast } from "react-toastify";
import loadingSpinner from "@/public/tube-spinner.svg"
import Image from "next/image";

interface Inputs {
    name: string;
    email: string;
    message: string;
}

export default function SupportSection() {
    const [loading, setLoading] = useState(false);

    const {
        register,
        handleSubmit,
        formState: { errors },
    } = useForm<Inputs>();

    const onSubmit: SubmitHandler<Inputs> = async (data) => {
        setLoading(true);

        toast.loading("Sending your message...", { toastId: "send-contact" });
        await axios
            .post("/api/contact", data)
            .then(() => {
                toast.update("send-contact", {
                    render: "Message sent! We'll get back to you ASAP 🤩",
                    type: "success",
                    isLoading: false,
                    autoClose: 6000,
                });
            })
            .catch((error) => {
                if (error.response.status === 400) {
                    toast.update("send-contact", {
                        render: error.response.data,
                        type: "error",
                        isLoading: false,
                        autoClose: 6000,
                    });
                } else {
                    toast.update("send-contact", {
                        render: "Oops! Something went wrong!\nI guess we need a developer that knows how to send emails 😗.\n You can always get in touch with us via Twitter, though.",
                        type: "error",
                        isLoading: false,
                        autoClose: 10000,
                    });
                }
            }).finally(() => setLoading(false));
    };

    return (
        <div className="mb-20 mx-4 md:mx-0" id="support">
            <div className="flex justify-center items-center">
                <h1 className=" text-3xl lg:text-4xl text-neutral-200 font-bold pb-6">
                    Contact us
                </h1>
            </div>
            <Card className="w-[100%] p-6 flex mx-auto bg-[#080908] border-2 border-[#161d15]">
                <form className="w-[100%]" onSubmit={handleSubmit(onSubmit)}>
                    <div className="flex-wrap md:flex-nowrap md:flex md:gap-4">
                        <input
                            type="text"
                            placeholder="Your Full Name @ Company XYZ"
                            className="focus:outline-none w-full p-2 mt-4 bg-[#0e0f0e] border border-[#161d15] rounded-md disabled:opacity-50 disabled:cursor-not-allowed"
                            {...register("name", { required: true })}
                            disabled={loading}
                        />
                        {errors.name && (
                            <span className="text-red-500">This field is required</span>
                        )}

                        <input
                            type="text"
                            placeholder="Email"
                            className="focus:outline-none w-full p-2 mt-4 bg-[#0e0f0e] border border-[#161d15] rounded-md disabled:opacity-50 disabled:cursor-not-allowed"
                            {...register("email", { required: true })}
                            disabled={loading}
                        />
                        {errors.email && (
                            <span className="text-red-500">
                                This field is required, we need it to contact you back
                            </span>
                        )}
                    </div>
                    <textarea
                        placeholder="Please type here the subject you wish to discuss with us..."
                        className="focus:outline-none w-full h-56 md:h-44 p-2 mt-4 bg-[#0e0f0e] border border-[#161d15] rounded-md disabled:opacity-50 disabled:cursor-not-allowed"
                        {...register("message", { required: true })}
                        disabled={loading}

                    />
                    {errors.message && (
                        <span className="text-red-500">This field is required</span>
                    )}

                    <Button
                        className={`
                            mt-2 px-12 
                            ${loading ? "bg-[#161d15] text-gray-200 border-2 border-gray-200" : "bg-primary"} 
                            ${loading ? "cursor-not-allowed" : "cursor-pointer"}
                        `}
                        type="submit"
                    >
                        {loading ?
                            <span className="flex items-center justify-center">
                                <Image src={loadingSpinner} alt="loading spinner" width={20} height={20} className="mr-2" />
                                Sending...
                            </span> : "Send"
                        }
                    </Button>
                </form>
                <div className="w-[40%] p-4 text-end hidden md:block">
                    <p className="text-3xl font-medium">Keep in touch, we</p>
                    <p className="text-3xl font-medium">want hear you!</p>
                </div>
            </Card>
        </div>
    );
}
