import React from 'react';

const ModalNote = ({active, setActive, children}) => {
    return (
        <div className={active ? "modal active" : "modal"} onClick={() => setActive(false)}>
            <div className="modal__content" onClick={e => e.stopPropagation()}>
                {children}
            </div>
        </div>
    );
};

export default ModalNote;