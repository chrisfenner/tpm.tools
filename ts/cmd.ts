import {
  GetAllCommandNamesResponse,
  CommandDescription,
  GetCommandDescriptionResponse,
} from "../proto/cmd";
import { cmdList, cmdLookup } from "./rest";

export async function lookupCommands() {
  const resultsElement = document.getElementById(
    "cmd-list-results"
  ) as HTMLDivElement;

  try {
    const rsp = await cmdList({});

    resultsElement.textContent = "";

    rsp.name.forEach((cmdName: string) => {
      var newA = document.createElement("a");
      newA.setAttribute("href", "/cmd/" + cmdName);
      var newDiv = document.createElement("div");
      newDiv.setAttribute("box-", "round");
      newDiv.innerText = cmdName;
      newA.appendChild(newDiv);
      resultsElement.appendChild(newA);
    });
  } catch (error) {
    console.error("Error:", error);
    alert("Failed to post data.");
  }
}
