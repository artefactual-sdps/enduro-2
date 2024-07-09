import PackageListLegend from "@/components/PackageListLegend.vue";
import { render, fireEvent, cleanup } from "@testing-library/vue";
import { describe, it, expect, afterEach } from "vitest";

describe("PackageListLegend.vue", () => {
  afterEach(() => cleanup());

  it("renders", async () => {
    const { getByText, getByLabelText, queryByText, emitted, rerender } =
      render(PackageListLegend, {
        props: {
          modelValue: true,
        },
      });

    getByText("DONE");
    getByText("ERROR");
    getByText("IN PROGRESS");
    getByText("QUEUED");
    getByText("PENDING");

    // Closing emits an event.
    const button = getByLabelText("Close");
    await fireEvent.click(button);
    expect(emitted()["update:modelValue"][0]).toStrictEqual([false]);

    // And setting the prop to false should hide the legend.
    await rerender({ modelValue: false });
    expect(queryByText("DONE")).toBeNull();
  });
});
