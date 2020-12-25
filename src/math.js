
class Compute {
    constructor() {
    }
    commandEval(message, args) {
        switch(args[0]){
            case 'add':
                let total = this.simpleAdd(args);
                (total.check)? message.reply(`Please check that you have proper spacing between your numbers`): message.reply(`The result is ${total.result}`);
                break;
            case 'minus':
                let final = this.simpleMinus(args);
                (final.check)? message.reply(`Please check that you have proper spacing between your numbers`): message.reply(`The result is ${final.result}`);
                console.log(final.check)
                break;
            case 'heads':
                let side = this.headsTails();
                message.reply(`You got ${side}`);
                console.log(`Coin flipped, got ${side}`);
                break;
        }
    };
    simpleAdd(nums) {
        let sum = 0;
        let isError = false;
        let res = {};
        nums.forEach((item, index, arr) => {
            try {
                if ((arr.indexOf(item) % 2 !== 0) && Number(item) !== NaN) {
                    sum += Number(item);
                };
            } catch(err) {
                console.log(`Error in Addition operation ${err}`);
                isError = true;
            }
        });
        if (sum === NaN) {isError = true};
        res.result = sum;
        res.check = isError;
        return res;
    };
    simpleMinus(nums) {
        let difference = Number(nums[1]);
        let isError = false;
        let res = {};
        nums.forEach((item, index, arr) => {
            try {
                if ((arr.indexOf(item) % 2 !== 0) && Number(item) !== NaN && arr.indexOf(item) != 1) {
                    difference -= Number(item);
                };
            } catch(err) {
                console.log(`Error in Subtraction operation ${err}`);
                isError = true;
            }
        });
        if (difference === NaN) {isError = true};
        console.log(difference);
        res.result = difference;
        res.check = isError;
        return res;
    };
    headsTails() {
        let evaluator = Math.floor(Math.random() * 11);
        let side = (evaluator % 2 == 0) ? 'Heads': 'Tails';
        return side;
    };
};
module.exports = {
    calc: new Compute()
}