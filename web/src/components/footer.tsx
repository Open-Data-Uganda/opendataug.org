import Link from "next/link";
import React from "react";
import CTA from "./cta";

interface FooterColumnProps {
  title: string;
  links: {
    label: string;
    href: string;
  }[];
}

const FooterColumn: React.FC<FooterColumnProps> = ({ title, links }) => {
  return (
    <div className="flex flex-col gap-4">
      <h3 className="font-semibold">{title}</h3>
      <ul className="flex flex-col gap-3">
        {links.map((link, index) => (
          <li key={index}>
            {link.href.startsWith("http") || link.href.startsWith("mailto") ? (
              <a
                href={link.href}
                className="text-base-content/70 hover:text-base-content transition-colors"
              >
                {link.label}
              </a>
            ) : (
              <Link
                href={link.href}
                className="text-base-content/70 hover:text-base-content transition-colors"
              >
                {link.label}
              </Link>
            )}
          </li>
        ))}
      </ul>
    </div>
  );
};

const Footer: React.FC = () => {
  const footerSections = [
    {
      title: "Company",
      links: [
        { label: "About us", href: "/about" },
        { label: "Careers", href: "/careers" },
        { label: "Blog", href: "/blog" },
      ],
    },
    {
      title: "Products",
      links: [
        { label: "Direct charge", href: "/direct-charge" },
        { label: "Checkout", href: "/checkout" },
        { label: "Payment links", href: "/payment-links" },
      ],
    },
    {
      title: "Resources",
      links: [
        { label: "Pricing", href: "/pricing" },
        { label: "Support", href: "/support" },
        { label: "Live businesses", href: "/live-businesses" },
        { label: "FAQs", href: "/faqs" },
      ],
    },
    {
      title: "Developers",
      links: [
        { label: "Community", href: "/community" },
        { label: "API Documentation", href: "/api-docs" },
        { label: "Status", href: "/status" },
      ],
    },
    {
      title: "Legal",
      links: [
        { label: "Terms of service", href: "/terms" },
        { label: "Privacy policy", href: "/privacy" },
        { label: "Cookie policy", href: "/cookie-policy" },
        { label: "End user agreement", href: "/eua" },
        { label: "AML statement", href: "/aml" },
      ],
    },
  ];

  return (
    <>
      <CTA />
      <footer className="border-base-content/10 border-b py-20">
        <div className="mx-auto">
          <div className="grid grid-cols-2 gap-8 md:grid-cols-3 lg:grid-cols-6">
            {/* Logo Column */}
            <div className="col-span-2 flex flex-col gap-4 md:col-span-3 lg:col-span-1">
              <Link href="/" className="flex items-center">
                <span className="text-2xl font-bold">ugandandata</span>
              </Link>
              <div className="flex items-center gap-2">
                <a
                  href="mailto:support@ugandandata.co"
                  className="text-base-content/70 hover:text-base-content"
                >
                  support@ugandandata.co
                </a>
              </div>
              <div className="flex gap-4">
                <a
                  href="https://facebook.com"
                  className="text-base-content/70 hover:text-base-content"
                  target="_blank"
                  rel="noopener noreferrer"
                >
                  <span className="sr-only">Facebook</span>
                  {/* Add Facebook Icon */}
                </a>
                {/* Add other social media links */}
              </div>
            </div>

            {/* Footer Columns */}
            {footerSections.map((section, index) => (
              <FooterColumn
                key={index}
                title={section.title}
                links={section.links}
              />
            ))}
          </div>
        </div>
      </footer>
      <div className=" border-gray-900/10 py-10">
        <p className="text-start text-sm text-gray-900">
          ©{new Date().getFullYear()}, Terra Con Technologies. All rights
          reserved
        </p>
      </div>
    </>
  );
};

export default Footer;
