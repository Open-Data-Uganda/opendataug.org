<script lang="ts">
  import CTA from "./CTA.svelte";

  interface Link {
    label: string;
    href: string;
  }

  interface FooterSection {
    title: string;
    links: Link[];
  }

   const FooterColumn = ({ title, links }: FooterSection) => {
    return {
      title,
      links
    };
  };

  const footerSections: FooterSection[] = [
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

  const currentYear = new Date().getFullYear();
</script>

<CTA />
<footer class="border-base-content/10 border-b py-20">
  <div class="mx-auto">
    <div class="grid grid-cols-2 gap-8 md:grid-cols-3 lg:grid-cols-6">
      <div class="col-span-2 flex flex-col gap-4 md:col-span-3 lg:col-span-1">
        <a href="/" class="flex items-center">
          <span class="text-2xl font-bold">opendataug</span>
        </a>
        <div class="flex items-center gap-2">
          <a
            href="mailto:support@opendataug.co"
            class="text-base-content/70 hover:text-base-content"
          >
            support@opendataug.co
          </a>
        </div>
        <div class="flex gap-4">
          <a
            href="https://facebook.com"
            class="text-base-content/70 hover:text-base-content"
            target="_blank"
            rel="noopener noreferrer"
          >
            <span class="sr-only">Facebook</span>
            <!-- Add Facebook Icon -->
          </a>
          <!-- Add other social media links -->
        </div>
      </div>

      {#each footerSections as section}
        <div class="flex flex-col gap-4">
          <h3 class="font-semibold">{section.title}</h3>
          <ul class="flex flex-col gap-3">
            {#each section.links as link}
              <li>
                {#if link.href.startsWith('http') || link.href.startsWith('mailto')}
                  <a
                    href={link.href}
                    class="text-base-content/70 hover:text-base-content transition-colors"
                  >
                    {link.label}
                  </a>
                {:else}
                  <a
                    href={link.href}
                    class="text-base-content/70 hover:text-base-content transition-colors"
                  >
                    {link.label}
                  </a>
                {/if}
              </li>
            {/each}
          </ul>
        </div>
      {/each}
    </div>
  </div>
</footer>


<div class="border-gray-900/10 py-10">
  <p class="text-start text-sm text-gray-900">
    ©{currentYear}, Terra Consults Technologies. All rights reserved
  </p>
</div>
