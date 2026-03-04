import { render, screen } from "@testing-library/react";
import { App } from "./App";

describe("App", () => {
  it("renders the heading", () => {
    render(<App />);
    expect(
      screen.getByRole("heading", { name: "PacketDeck" }),
    ).toBeInTheDocument();
  });

  it("renders the description", () => {
    render(<App />);
    expect(
      screen.getByText("Container network topology visualization"),
    ).toBeInTheDocument();
  });
});
