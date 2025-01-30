import Link from "next/link";
import React from "react";

const CTA: React.FC = () => {
  return (
    <div className="mx-auto my-16 bg-primary px-4 py-16 text-white lg:rounded-lg lg:px-10">
      <div className="flex flex-col justify-center gap-8">
        <h2 className="mx-auto max-w-xl text-4xl font-bold leading-tight">
          Ready to integrate Uganda's comprehensive location data into your
          applications?
        </h2>
        <Link
          href="/signup"
          className="mx-auto w-fit rounded-lg bg-white px-6 py-3 font-semibold text-primary transition-colors hover:bg-white/90"
        >
          Get your API key
        </Link>
      </div>
    </div>
  );
};

export default CTA;
