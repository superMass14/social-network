export const ShowFollowings = ({ followings, isVisible, onClose }) => {
  if (!isVisible) return null;
  return (
    <div
      className="box-border fixed inset-0 bg-bg bg-opacity-10 backdrop-blur-sm
        flex justify-center items-center"
    >
      <div className="w-1/3 h-[600px] rounded-lg border bg-bg p-4 shadow-md sm:p-8">
        <div className="mb-4 flex items-center justify-between">
          <h3 className="text-xl font-bold leading-none">Followings</h3>
          <button className="w-full flex justify-end" onClick={() => onClose()}>
            <svg
              xmlns="http://www.w3.org/2000/svg"
              viewBox="0 0 24 24"
              fill="currentColor"
              className="w-8 h-8 hover:text-red-500
                   hover:rotate-90 transition duration-300 ease-in-out place-self-end"
            >
              <path
                fillRule="evenodd"
                d="M5.47 5.47a.75.75 0 0 1 1.06 0L12 10.94l5.47-5.47a.75.75 0 1 1 1.06 1.06L13.06 12l5.47 5.47a.75.75 0 1 1-1.06 1.06L12 13.06l-5.47 5.47a.75.75 0 0 1-1.06-1.06L10.94 12 5.47 6.53a.75.75 0 0 1 0-1.06Z"
                clipRule="evenodd"
              />
            </svg>
          </button>
        </div>
        <div className="flow-root h-[90%]">
          <ul
            role="list"
            className="h-full divide-y divide-gray-900 overflow-y-auto"
          >
            {followings
              ? followings.map((following) => followingCard(following))
              : "No followings found."}
          </ul>
        </div>
      </div>
    </div>
  );
};

export const followingCard = (following) => {
  return (
    <li key={following.email} className="shadow-inner py-2.5 sm:py-3.5">
      <div className="flex items-center space-x-4">
        <div className="flex-shrink-0">
          <img
            className="h-12 w-12 rounded-full"
            src={following.avatar}
            alt={`${following.first_name} image`}
          />
        </div>
        <div className="min-w-0 flex-1">
          <p className="truncate text-sm font-medium">
            {`${following.first_name} ${following.last_name}`}
          </p>
          <p className="truncate text-sm text-gray-400">{following.email}</p>
        </div>
      </div>
    </li>
  );
};
