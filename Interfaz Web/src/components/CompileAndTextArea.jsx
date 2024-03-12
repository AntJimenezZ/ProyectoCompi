import React,{useState} from "react";
import {Button, Textarea} from "@nextui-org/react";

export default function App() {


  const [text, setText] = useState('');
  const [responseText, setResponseText] = useState(''); 

  const handleClick = async () => {
    try {
      const response = await fetch('http://localhost:8080/json', {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json'
        },
        body: JSON.stringify({ text: text })
      });

      const data = await response.json();
      setResponseText(data.text);
    } catch (error) {
      console.error('Error:', error);
    }
  };


  return (

    <div className="flex flex-col items-center space-y-4">

      <h1>
        Write your code here:
      </h1>
      <Textarea
        label="Code:"
        variant="flat"
        placeholder="You can make this textarea resizable by dragging the bottom right corner."
        disableAnimation
        //disableAutosize
        value={text}
        onChange={(e) => setText(e.target.value)}
        classNames={{
          base: "max-w-xs  ",
          input: "resize-y min-h-[270px]",
        }}
      />
      <h1>
        <Button color="success" variant="shadow" onPress={handleClick}>
          Compile
        </Button>
      </h1>
      
    </div>
    
  );
}
