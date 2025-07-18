import json
from web3 import Web3

contract_list = ["pool", "atoken", "vtoken", "gateway"]
signatures = []

for contract in contract_list:
    with open(f"abi/{contract}.abi") as f:
        abi = json.load(f)
    for element in abi:
        if element["type"] == "event":
            event_inputs = []
            for input in element["inputs"]:
                event_inputs.append(input["type"])
                # event_inputs.append(input["name"]) ## CHANGE HERE TO GET SIGNATURE
            event_signature = element["name"] #+ "(" + ",".join(event_inputs) + ")"
            signatures.append(event_signature)

signatures = list(set(signatures))

event_to_hash = {}

for signature in signatures:
    signature_hash = Web3.to_hex(Web3.keccak(text=signature))
    event_to_hash[signature] = signature_hash

for k, v in event_to_hash.items():
    # print('"'+k+'"', ":", '"'+v+'",')
    # print("case eventsCodes[", '"'+k+'"', "]:")
    print(f'datalab.SaveRecords(endpoint, accessKeyID, secretAccessKey, allDecodedEvents.{k}, bucket, output_path + "decoded_{k}.json")')
# print(event_to_hash)
