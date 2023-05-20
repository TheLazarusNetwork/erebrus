import React, { useState } from 'react';
import { createClient, CreateClientPayload } from '../../modules/api';

const AddClient = () => {
  const [name, setName] = useState("");
  const [email, setEmail] = useState('');
  const [tags, setTags] = useState('');
  const [enable, setEnable] = useState(true);
  const adminEmail = "adimis.ai.001@gmail.com"

  const handleAddClient = async (event: React.FormEvent<HTMLFormElement>) => {
    event.preventDefault();

    const tagsArray = tags.split(',').map(tag => tag.trim());

    const payload: CreateClientPayload = {
      name,
      tags: tagsArray,
      email,
      enable,
      allowedIPs: [
        "0.0.0.0/0",
        "::/0"
      ],
      address: [
        "10.0.0.1/24"
      ],
      createdBy: adminEmail
    };

    try {
      await createClient(payload);
      window.location.reload();
    } catch (error) {
    }
  };

  return (
    <div>
        <label htmlFor="add-client-modal">
        Add Client
        </label>
        <input type="checkbox" id="add-client-modal" className="modal-toggle" />
        <div className="modal">
            <div className="modal-box relative w-full bg-gray-900 rounded-md bg-clip-padding backdrop-filter backdrop-blur-sm bg-opacity-80 border border-gray-100">
            <label htmlFor="add-client-modal" className="btn btn-sm btn-circle absolute right-2 top-2">✕</label>
            <h2 className="font-bold text-blue-200 text-3xl">Add Client</h2>
                <div>
                    <form onSubmit={handleAddClient}>
                        <div className="space-y-4">
                            <div>
                                <label htmlFor="name" className="block text-blue-100 text-md text-left my-3 font-semibold">
                                Name
                                </label>
                                <input
                                id="name"
                                value={name}
                                onChange={(e) => setName(e.target.value)}
                                className="input input-bordered w-full max-w-x text-black"
                                />
                            </div>
                            <div>
                                <label htmlFor="email" className="block text-blue-100 text-md text-left my-3 font-semibold">
                                Email
                                </label>
                                <input
                                id="email"
                                value={email}
                                onChange={(e) => setEmail(e.target.value)}
                                className="input input-bordered w-full max-w-x text-black"
                                />
                            </div>
                            <div>
                                <label htmlFor="tags" className="block text-blue-100 text-md text-left my-3 font-semibold">
                                Tags (comma-separated)
                                </label>
                                <input
                                id="tags"
                                value={tags}
                                onChange={(e) => setTags(e.target.value)}
                                className="input input-bordered w-full max-w-x text-black"
                                />
                            </div>
                            <div className='form-control'>
                                <label className="label cursor-pointer">
                                <span className="label-text text-gray-100 text-md">Enable/Disable User</span> 
                                <input
                                id="enable"
                                type="checkbox"
                                checked={enable}
                                onChange={(e) => setEnable(e.target.checked)}
                                className="toggle-sm toggle toggle-info my-3"
                                />
                                </label>
                            </div>
                        </div>
                        <button type="submit" className="bg-gradient-to-r from-blue-300 to-blue-500 text-gray-900 font-semibold rounded-lg p-3 px-5 mt-4">
                            Submit
                        </button>
                    </form>
                </div>
            </div>
        </div>
    </div>
    );
    };
    
    export default AddClient;