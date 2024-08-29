export const abi = [
    {
        type: 'function',
        name: 'approve',
        stateMutability: 'nonpayable',
        inputs: [
            { name: 'spender', type: 'address' },
            { name: 'amount', type: 'uint256' },
        ],
        outputs: [{ type: 'bool' }],
    },

    {
        type: 'function',
        name: 'depositERC20Tokens',
        stateMutability: 'nonpayable',
        inputs: [
            {
                name: '_token',
                type: 'address'
            },
            {
                name: '_dapp',
                type: 'address'
            },
            {
                name: '_amount',
                type: 'uint256'
            },
            {
                name: '_execLayerData',
                type: 'bytes'
            }
        ],
        outputs: [],

    },
] as const


