import React, { useState, useEffect } from 'react';
import { marked } from 'marked';

const MentaAI: React.FC = () => {
    const [messages, setMessages] = useState<{ user: string, text: string, timestamp: string }[]>([]);
    const [input, setInput] = useState('');
    const [loading, setLoading] = useState(false);
    const [patientId, setPatientId] = useState<string | null>(null);

    useEffect(() => {
        setPatientId("67a932d0ca3e10e26a9d8c40"); // Replace with actual patient ID
    }, []);

    const handleSend = async () => {
        if (input.trim() && patientId) {
            const timestamp = new Date().toLocaleTimeString();
            setMessages([...messages, { user: 'You', text: input, timestamp }]);
            setLoading(true);

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
                setMessages((prevMessages) => [...prevMessages, { user: 'MentaAI', text: data.bot_response, timestamp: new Date().toLocaleTimeString() }]);
            } catch (error) {
                console.error('Error:', error);
                setMessages((prevMessages) => [...prevMessages, { user: 'MentaAI', text: 'Error connecting to the server.', timestamp: new Date().toLocaleTimeString() }]);
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

    const createMarkup = (text: string) => {
        return { __html: marked(text) };
    };

    return (
        <div className="flex flex-col h-screen p-4 bg-gray-100">
            <div className="flex-none p-4 bg-green-500 text-white rounded-t-lg flex items-center justify-between">
            <h2 className="text-xl font-bold">MentaAI Chat</h2>
            </div>
            <div className="flex-1 p-4 bg-white rounded-b-lg shadow-md overflow-y-auto">
            {messages.map((message, index) => (
                <div
                key={index}
                className={`mb-3 p-3 rounded-lg ${message.user === 'You' ? 'bg-green-100 self-end' : 'bg-gray-200 self-start'}`}
                >
                <strong>{message.user}:</strong> <span className="text-xs text-gray-500">{message.timestamp}</span>
                <div dangerouslySetInnerHTML={createMarkup(message.text)} />
                </div>
            ))}
            {loading && <p className="text-gray-500 italic">MentaAI is typing...</p>}
            </div>
            <div className="flex-none mt-4" style={{ height: '20%' }}>
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
                className="p-2 bg-green-500 text-white rounded-r-lg"
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