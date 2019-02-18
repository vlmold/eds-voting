pragma solidity 0.5.3;

contract VotingStorage {

    address owner;
    mapping (uint64 => address)  indexes ;
    mapping (address => Vote) votes ;
    uint64 votesCount = 0 ;

    mapping (address => bool) whiteList ;
    struct Vote {
        address voter;
        string encryptedOption;
        bytes signature;
    }
    constructor () public  {
        owner = msg.sender;
    }
    modifier onlyOwner () {
        require(msg.sender == owner);
        _;
    }
    function addToWhiteList(address  _voter) public onlyOwner {
        whiteList[_voter] = true;
    }
    function registerVote (string memory _vote , bytes memory _sig) public {
        votesCount++;
        bytes32 message = sigHash(abi.encodePacked(_vote, msg.sender));
        address addr = recoverSigner(message, _sig);
        require(addr == msg.sender);

        Vote memory newVote;
        newVote.encryptedOption = _vote;
        newVote.signature = _sig;
        newVote.voter = msg.sender;

        votes[msg.sender] = newVote;
    }
    function sigHash(bytes memory _message) private pure returns(bytes32){

        bytes32 preMessage = keccak256(_message);

        return prefixed(preMessage);
    }

    // Internal functions
    // ...
    // Builds a prefixed hash to mimic the behavior of eth_sign.
    function prefixed(bytes32 hash) internal pure returns (bytes32) {

        return  keccak256(
            abi.encodePacked("\x19Ethereum Signed Message:\n32", hash)
        );

    }
    function recoverSigner(bytes32 message, bytes memory sig)
    internal
    pure
    returns (address)
    {
        uint8 v;
        bytes32 r;
        bytes32 s;

        (v, r, s) = splitSignature(sig);

        address sr = ecrecover(message, v, r, s);

        return sr;
    }

    function splitSignature(bytes memory sig)
    internal
    pure
    returns (uint8, bytes32, bytes32){

        require(sig.length == 65);

        bytes32 r;
        bytes32 s;
        uint8 v;

        assembly {
        // first 32 bytes, after the length prefix
            r := mload(add(sig, 32))
        // second 32 bytes
            s := mload(add(sig, 64))
        // final byte (first byte of the next 32 bytes)
            v := byte(0, mload(add(sig, 96)))
        }

        return (v, r, s);
    }
}