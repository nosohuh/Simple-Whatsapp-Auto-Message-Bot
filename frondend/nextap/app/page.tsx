import { Snippet } from "@nextui-org/snippet";
import { Code } from "@nextui-org/code";

import { title, subtitle } from "@/components/primitives";

export default function Home() {
  return (
    <section className="flex flex-col items-center justify-center gap-4 py-8 md:py-10">
      <div className="inline-block max-w-xl text-center justify-center">
        <span className={title()}>Bu&nbsp;</span>
        <span className={title({ color: "violet" })}>Web Sitesi</span>
        <br />
        <span className={title()}>ülkeniz de desteklenmemektedir..</span>
        <div className={subtitle({ class: "mt-4" })}>---</div>
      </div>

      <div className="mt-8">
        <Snippet hideCopyButton hideSymbol variant="bordered">
          <span>
            <Code color="primary">Lütfen Çıkınız.</Code>
          </span>
        </Snippet>
      </div>
    </section>
  );
}

