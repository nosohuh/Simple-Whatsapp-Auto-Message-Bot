import React from 'react';
import { EditUserAction } from '../../axios/axios'; // API çağrısı yapılacak
import Noty from 'noty';
import 'noty/lib/noty.css'; // veya `import 'noty';`
import 'noty/lib/themes/mint.css'; // tema

const DeleteButton = ({ id }) => {
  const handleDelete = async () => {
    // Noty ile evet/hayır sorusu
    const noty = new Noty({
      text: 'Bu kullanıcıyı silmek istediğinize emin misiniz?',
      type: 'alert',
      layout: 'center',
      buttons: [
        Noty.button('Evet', 'btn btn-success mr-10', async function () {
          try {
            const request = await EditUserAction(
              'delete',
              null,
              null,
              id,
              null,
              null,
              null
            );
            noty.close();
            window.location.reload();
          } catch (error) {
            if (error.response.status === 401) {
              new Noty({
                type: 'error',
                text: "Yetkisiz erişim!",
                timeout: 3000,
                theme: 'mint',
              }).show();
            }
            noty.close();
          }
        }),
        Noty.button('Hayır', 'btn btn-danger', function () {
          noty.close();
        })
      ]
    }).show();
  };


  // DeleteIcon Component as a part of DeleteButton
  const DeleteIcon = (props) => (
    <svg
      aria-hidden="true"
      fill="none"
      focusable="false"
      height="1em"
      role="presentation"
      viewBox="0 0 20 20"
      width="1em"
      {...props}
    >
      <path
        d="M17.5 4.98332C14.725 4.70832 11.9333 4.56665 9.15 4.56665C7.5 4.56665 5.85 4.64998 4.2 4.81665L2.5 4.98332"
        stroke="currentColor"
        strokeLinecap="round"
        strokeLinejoin="round"
        strokeWidth={1.5}
      />
      <path
        d="M7.08331 4.14169L7.26665 3.05002C7.39998 2.25835 7.49998 1.66669 8.90831 1.66669H11.0916C12.5 1.66669 12.6083 2.29169 12.7333 3.05835L12.9166 4.14169"
        stroke="currentColor"
        strokeLinecap="round"
        strokeLinejoin="round"
        strokeWidth={1.5}
      />
      <path
        d="M15.7084 7.61664L15.1667 16.0083C15.075 17.3166 15 18.3333 12.675 18.3333H7.32502C5.00002 18.3333 4.92502 17.3166 4.83335 16.0083L4.29169 7.61664"
        stroke="currentColor"
        strokeLinecap="round"
        strokeLinejoin="round"
        strokeWidth={1.5}
      />
      <path
        d="M8.60834 13.75H11.3833"
        stroke="currentColor"
        strokeLinecap="round"
        strokeLinejoin="round"
        strokeWidth={1.5}
      />
      <path
        d="M7.91669 10.4167H12.0834"
        stroke="currentColor"
        strokeLinecap="round"
        strokeLinejoin="round"
        strokeWidth={1.5}
      />
    </svg>
  );

  return (
    <DeleteIcon onClick={handleDelete} style={{ cursor: 'pointer' }} />
  );
};

export default DeleteButton;
