import React, { useState, useEffect } from 'react';
import { marked } from 'marked'; // Import the marked library

const MentaAI: React.FC = () => {
    const [messages, setMessages] = useState<{ user: string, text: string }[]>([]);
    const [input, setInput] = useState('');
    const [loading, setLoading] = useState(false);
    const [patientId, setPatientId] = useState<string | null>(null);

    useEffect(() => {
        // Fetch or set patient_id dynamically (this can be modified to fetch from a login system)
        setPatientId("67a932d0ca3e10e26a9d8c40"); // Replace with actual patient ID
    }, []);

    const handleSend = async () => {
        if (input.trim() && patientId) {
            setMessages([...messages, { user: 'You', text: input }]);
            setLoading(true); // Show loading

            try {
                const response = await fetch('http://localhost:8000/chat/chatWithBot', {
                    method: 'POST',
                    headers: { 'Content-Type': 'application/json' },
                    body: JSON.stringify({
                        patient_id: patientId,
                        user_message: input,
                    }),
                });

                const data = await response.json();
                setMessages((prevMessages) => [...prevMessages, { user: 'MentaAI', text: data.bot_response }]);
            } catch (error) {
                console.error('Error:', error);
                setMessages((prevMessages) => [...prevMessages, { user: 'MentaAI', text: 'Error connecting to the server.' }]);
            }

            setLoading(false);
            setInput('');
        }
    };

    const handleKeyPress = (e: React.KeyboardEvent) => {
        if (e.key === 'Enter' && !loading) {
            handleSend();
        }
    };

    // Function to convert markdown text to HTML using `marked`
    const createMarkup = (text: string) => {
        const htmlContent = marked(text);  // Convert markdown text to HTML
        return { __html: htmlContent };
    };

    return (
        <div className="flex flex-col h-screen p-4 bg-gray-100">
            <div className="flex-none p-4 bg-blue-500 text-white rounded-t-lg">
                <h2 className="text-xl font-bold">MentaAI Chat</h2>
            </div>
            <div className="flex-1 overflow-y-auto p-4 bg-white rounded-b-lg shadow-md">
                {messages.map((message, index) => (
                    <div
                        key={index}
                        className={`mb-3 p-3 rounded-lg ${message.user === 'You' ? 'bg-blue-100 self-end' : 'bg-gray-200 self-start'}`}
                    >
                        <strong>{message.user}:</strong>
                        <div
                            dangerouslySetInnerHTML={createMarkup(message.text)}  // Use `dangerouslySetInnerHTML` to render HTML
                        />
                    </div>
                ))}
                {loading && <p className="text-gray-500 italic">MentaAI is thinking...</p>}
            </div>
            <div className="flex-none mt-4">
                <div className="flex">
                    <input
                        type="text"
                        value={input}
                        onChange={(e) => setInput(e.target.value)}
                        onKeyDown={handleKeyPress}
                        placeholder="Type your message..."
                        className="flex-1 p-2 border rounded-l-lg"
                    />
                    <button
                        onClick={handleSend}
                        className="p-2 bg-blue-500 text-white rounded-r-lg"
                        disabled={loading}
                    >
                        {loading ? "Sending..." : "Send"}
                    </button>
                </div>
            </div>
        </div>
    );
};

export default MentaAI;
