# Braindump
Currently getting user input from the state of the model. No clue how this will change with additional features to model. 

### Application 
- singular model? or multiple models that trigger with events?
- clearing users input at event enter? 
    - reset textInput state to original value at declaration in initialModel?
- copy textInput state at enter event, write to []string, at tea.Quit case also call save() to write json

### View
- currently redisplaying formatted text in View() at every enter event
    - Sprintf is returned everytime model is updated -> Lipgloss to format the display that returns blank or resets the state? maybe just don't want to print new line?
- research more into FancyList for View()

