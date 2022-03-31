let ResetOptionsAfterEachSubmit = true;

$(document).ready(function() {   
    initialize();
});

function initialize() {
    var numbers = [0,1,2,3,4,5,6,7,8,9,10,11,12];
    for (let index = 1; index <= 5; index++) {
        var selectTag = document.getElementById("val"+index);
        numbers.forEach(item => {
            var option = document.createElement("option");
            option.value = item;
            option.text = item;
            if (item == 0) {
                option.selected = true
            }
            selectTag.add(option)
        });
    }
}


function submit() {
    var selectedNumbers = [];
    for (let index = 1; index <= 5; index++) {
        var selectTag = document.getElementById("val"+index).value;
        if (!isInArray(selectTag, selectedNumbers)) {
            selectedNumbers.push(selectTag);   
        } else {
            alert("It only needs a little bit of brain to not select same options for this game!");
            return
        }
    }

    var guessedNumbers = {
        GuessedNumberOne: parseInt(document.getElementById("val1").value),
        GuessedNumberTwo: parseInt(document.getElementById("val2").value),
        GuessedNumberThree: parseInt(document.getElementById("val3").value),
        GuessedNumberFour: parseInt(document.getElementById("val4").value),
        GuessedNumberFive: parseInt(document.getElementById("val5").value)
    };
    guessedNumbers = JSON.stringify(guessedNumbers)
    fetch("/ReceiveInput?a=" + guessedNumbers)
    .then(response => response.json())
    .then(data => { 
        if (ResetOptionsAfterEachSubmit) {
            resetToDefault();
        } else {
            resetResult();
        }
        document.getElementById("generatedNumbers").innerHTML = "We generated the numbers<br>" + data.generated_numbers;
        document.getElementById("guessedNumbers").innerHTML = "You guessed the numbers<br>" + data.guessed_numbers;
        if (data.match_counter == 0) {
            document.getElementById("result").innerHTML =   "<br><b>Result :</b><br>" + 
                                                            "You guessed none of the numbers we generated!<br><br>" + 
                                                            "<div class=\"container-red m-t-20\">" + 
                                                            "<b>You’re an APPRENTICE guesser!</b><br>" + 
                                                            "</div>";
        } else if (data.match_counter < 5) {
            document.getElementById("result").innerHTML =   "<br><b>Result :</b><br>" + 
                                                            "You guessed " + data.match_counter + " of the numbers we generated!<br><br>" + 
                                                            "<div class=\"container-blue m-t-20\">" + 
                                                            "<b>You’re a GOOD guesser!</b><br>" + 
                                                            "</div>";
        } else if (data.match_counter == 5) {
            document.getElementById("result").innerHTML =   "<br><b>Result :</b><br>" + 
                                                            "You guessed all of of the numbers we generated!<br><br>" + 
                                                            "<div class=\"container-green m-t-20\">" + 
                                                            "<b>You’re an EXCELLENT guesser!</b><br>" + 
                                                            "</div>";
        }
    })
}

function resetToDefault() {
    resetResult();
    resetOptions();
}

function resetResult() {
    document.getElementById("generatedNumbers").innerHTML = "\u00A0"; //$nbsp;
    document.getElementById("guessedNumbers").innerHTML = "\u00A0"; //$nbsp;
    document.getElementById("result").innerHTML = "\u00A0"; //$nbsp;
}

function resetOptions() {
    for (let index = 1; index <= 5; index++) {
        var selectTag = document.getElementById("val"+index);
        selectTag.selectedIndex = 0;
    }
}


// TODO - Block Already Selected Options From Other Select Tags Instead Of Calling This Function
function isInArray(element, array) {
    for(let i = 0; i < array.length; i++) {
        if (array[i] == element) {
            return true;
        }
    }
    return false;
}